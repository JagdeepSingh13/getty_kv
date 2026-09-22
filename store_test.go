package main

import (
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
