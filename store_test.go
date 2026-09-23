package main

import (
	"errors"
	"reflect"
	"testing"
)

// using the testing package and TestXxx method

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("charlie", "3")
	store.Set("bravo", "2")
	store.Set("alpha", "1")

	got := store.Keys()
	want := []string{"alpha", "bravo", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want = %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()
	// store.Set("a", "45")

	got := store.Keys()
	if len(got) != 0 {
		t.Errorf("Keys() on empty store = %v, want empty slice", got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()
	store.Set("hello", "world")

	// should give "world", nil
	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("Get() failed")
	}

	// should give "", err
	if val, err := store.Get("missing"); err == nil || val != "" {
		t.Errorf("Get() failed")
	}
}

func TestSetGet_EmptyKeys(t *testing.T) {
	store := NewStore()

	if _, err := store.Get(""); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Error("Get() failed")
	}
}
