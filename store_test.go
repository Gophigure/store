package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	store := new(Store[int, int])

	store.Set(1, 1)

	if one, ok := store.Get(1); !ok {
		t.Error("expected key \"one\" with value 1 to be present but was not")
		return
	} else if one != 1 {
		t.Errorf("got %v but 1 was expected", one)
		return
	} else {
		t.Logf("got %v for key \"one\"", one)
	}

	store.Set(2, 2)

	store.ForEach(func(key int, value int) {
		t.Logf("got %v for key \"%v\"", value, key)
	})

	store.Delete(1)

	if _, ok := store.Get(1); ok {
		t.Error("expected key \"one\" to not be present")
		return
	} else {
		t.Log("confirmed key \"one\" deleted successfully")
	}

	if two, ok := store.Pluck(2); !ok {
		t.Error("expected key \"two\" to be present")
		return
	} else if two != 2 {
		t.Errorf("expected plucked key \"two\" to equal 2 but got %v", two)
		return
	} else {
		t.Logf("retrieved key \"two\" with value %v with Pluck", 2)
	}

	_, ok := store.Get(2)
	if ok {
		t.Error("expected key \"two\" to not be present")
		return
	} else {
		t.Log("confirmed key \"two\" plucked successfully")
	}

	for i := 1; i <= 10; i++ {
		store.Set(i, i)
	}

	store.ForEach(func(key int, value int) {
		t.Logf("got %v for key \"%v\"", value, key)
	})

	t.Log("resetting data")
	store.Reset()

	store.ForEach(func(_ int, _ int) {
		t.Error("store should have no keys to iterate")
		return
	})

	t.Log("successfully reset data")
}
