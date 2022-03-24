package store

import (
	"strconv"
	"testing"
)

func TestStore(t *testing.T) {
	store := new(Store[string, int64])

	store.Set("one", 1)

	if one, ok := store.Get("one"); !ok {
		t.Error("expected key \"one\" with value 1 to be present but was not")
	} else if one != 1 {
		t.Errorf("got %v but 1 was expected", one)
	} else {
		t.Logf("got %v for key \"one\"", one)
	}

	store.Set("two", 2)

	store.ForEach(func(key string, value int64) {
		t.Logf("got %v for key \"%v\"", value, key)
	})

	store.Delete("one")

	if _, ok := store.Get("one"); ok {
		t.Error("expected key \"one\" to not be present")
	} else {
		t.Log("confirmed key \"one\" deleted successfully")
	}

	if two, ok := store.Pluck("two"); !ok {
		t.Error("expected key \"two\" to be present")
	} else if two != 2 {
		t.Errorf("expected plucked key \"two\" to equal 2 but got %v", two)
	} else {
		t.Logf("retrieved key \"two\" with value %v with Pluck", 2)
	}

	_, ok := store.Get("two")
	if ok {
		t.Error("expected key \"two\" to not be present")
	} else {
		t.Log("confirmed key \"two\" plucked successfully")
	}

	for i := int64(1); i <= 10; i++ {
		store.Set(strconv.FormatInt(i, 10), i)
	}

	store.ForEach(func(key string, value int64) {
		t.Logf("got %v for key \"%v\"", value, key)
	})

	t.Log("resetting data")
	store.Reset()

	store.ForEach(func(_ string, _ int64) {
		t.Error("store should have no keys to iterate")
	})

	t.Log("successfully reset data")
}
