package main

import (
	"errors"
	"sort"
)

var ErrKeyDoesNotExists = errors.New("key does not exists")

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// make the test, run it first

func (s *Store) Rename(oldKey, newKey string) {
	// remove the val and oldKey, insert val with newKey
	if val, err := s.Get(oldKey); err != nil {
		s.Delete(oldKey)
		s.Set(newKey, val)
	}
}

func (s *Store) Pop(key string) (string, bool) {
	val, _ := s.Get(key)

	s.Delete(key)
	return val, true
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	// make slice of keys, sort the keys alb.
	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExists
	}

	return val, nil
}

func (s *Store) Set(key, val string) {
	s.data[key] = val
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Len() int {
	return len(s.data)
}
