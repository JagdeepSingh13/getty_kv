package main

import "fmt"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, val string) {
	s.data[key] = val
}

func main() {
	s := NewStore()

	s.Set("a", "45")
	s.Set("b", "50")

	fmt.Println(s.Get("a"))
}
