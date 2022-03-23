package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	store := new(Store[string, int])

	store.Set("one", 1)

	if one, ok := store.Get("one"); !ok {
		t.Errorf("expected key \"one\" with value 1 to be present but was not")
	} else if one != 1 {
		t.Errorf("got %v but 1 was expected", one)
	} else {
		t.Logf("got %v for key \"one\"", one)
	}

	store.Set("two", 2)

	store.ForEach(func(key string, value int) {
		t.Logf("got %v for key \"%v\"", value, key)
	})

	store.Delete("one")

	if _, ok := store.Get("one"); ok {
		t.Errorf("expected one to not be present")
	}
}
