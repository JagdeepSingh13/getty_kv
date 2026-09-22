package main

import (
	"fmt"
	"sort"
)

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// make the test, run it first

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	// make slice of keys, sort the keys alb.
	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, val string) {
	s.data[key] = val
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func main() {
	s := NewStore()

	s.Set("a", "45")
	s.Set("b", "50")

	fmt.Println(s.Get("a"))
}
