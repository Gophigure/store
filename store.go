// Package store provides Store, which is an alternative to the sync.Map type
// and supports both type-safety using 1.18 generics and avoids the unsafe
// package.
package store

import (
	"sync"
	"sync/atomic"
)

// Store is a go-routine safe way to store an arbitrary map. Its zero-value is
// safe to use. It must not be copied after first use.
type Store[T comparable, K any] struct {
	_      [0]func() // prevent Store{} & comparison
	mu     sync.Mutex
	clean  atomic.Value
	stale  bool
	dirty  map[T]*K
	misses int
}

// ensureClean simply ensures that Store.clean is valid, and if it isn't, make
// it so.
func (s *Store[T, K]) ensureClean() (clean map[T]*K) {
	if loaded := s.clean.Load(); loaded != nil {
		clean = loaded.(map[T]*K)
	} else {
		clean = make(map[T]*K, 1)
		s.clean.Store(clean)
	}
	return
}

// Set allows you to store value under key.
func (s *Store[T, K]) Set(key T, value K) {
	clean := s.ensureClean()
	if _, ok := clean[key]; ok {
		clean[key] = &value
		return
	}

	s.mu.Lock()
	clean = s.clean.Load().(map[T]*K)
	if val, ok := clean[key]; ok {
		s.dirty[key], clean[key] = val, &value
	} else if _, ok = s.dirty[key]; ok {
		s.dirty[key] = &value
	} else {
		if !s.stale && s.dirty == nil {
			s.dirty, s.stale = make(map[T]*K, len(clean)), true
			s.clean.Store(s.dirty)
		}
		s.dirty[key] = &value
	}
	s.mu.Unlock()
}

// GetOrSet only stores value under key if no value already exists under key,
// returns true if successful.
func (s *Store[T, K]) GetOrSet(key T, value K) (get K, set bool) {
	clean := s.ensureClean()
	if v, ok := clean[key]; ok {
		get = *v
		return
	}

	s.mu.Lock()
	if v, ok := clean[key]; ok {
		get, s.dirty[key] = *v, v
	} else if v, ok = s.dirty[key]; ok {
		get = *v
		s.misses++
	} else {
		if !s.stale && s.dirty == nil {
			s.dirty, s.stale = make(map[T]*K, len(clean)), true
			for k, val := range clean {
				s.dirty[k] = val
			}
			s.clean.Store(clean)
		}
		s.dirty[key], get, set = &value, value, false
	}

	s.mu.Unlock()

	return
}

// Get returns a value from the Store, if it exists.
func (s *Store[T, K]) Get(key T) (K, bool) {
	clean := s.ensureClean()
	if val, ok := clean[key]; ok {
		return *val, ok
	} else if !ok && s.stale {
		s.mu.Lock()
		clean = s.clean.Load().(map[T]*K)
		val, ok = clean[key]
		// re-check after acquiring lock
		if !ok && s.stale {
			val, ok = s.dirty[key]
			s.misses++
			if s.misses >= len(s.dirty) {
				s.clean.Store(s.dirty)
				s.dirty, s.misses = make(map[T]*K, 1), 0
			}
		}
		s.mu.Unlock()
	}

	return *new(K), false
}

// Pluck both retrieves and removes a key from the Store.
func (s *Store[T, K]) Pluck(key T) (K, bool) {
	clean := s.ensureClean()
	if val, ok := clean[key]; ok {
		delete(clean, key)
		return *val, true
	} else if !ok && s.stale {
		s.mu.Lock()
		val, ok = clean[key]
		if !ok && s.stale {
			val = s.dirty[key]
			delete(s.dirty, key)
			s.misses++
			if s.misses >= len(s.dirty) {
				s.clean.Store(s.dirty)
				s.dirty, s.misses = make(map[T]*K, 1), 0
			}
		}
		return *val, ok
	}

	return *new(K), false
}

// Delete removes a key from the Store, no-op if no value is held.
func (s *Store[T, K]) Delete(key T) { _, _ = s.Pluck(key) }

// ForEach calls the provided function for each value stored.
func (s *Store[T, K]) ForEach(f func(key T, value K)) {
	if !s.stale {
		for key, value := range s.ensureClean() {
			f(key, *value)
		}
		return
	}

	s.mu.Lock()
	for key, value := range s.dirty {
		f(key, *value)
	}
	s.mu.Unlock()
}

// Reset removes all stored data from the Store.
func (s *Store[T, K]) Reset() {
	s.mu.Lock()
	m := make(map[T]*K)
	s.clean.Store(m)
	s.dirty, s.stale, s.misses = m, false, 0
	s.mu.Unlock()
}
