package main

import (
	"errors"
	"fmt"
	"sort"
)

var ErrKeyDoesNotExists = errors.New("key does not exists")
var ErrEmptyKey = errors.New("key is mandatory")
var ErrorStoreFull = errors.New("store is full")

type Store struct {
	data    map[string]string
	maxSize int
}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize, // 0 -> unlimited
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
	s.maxSize -= 1
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
	if key == "" {
		return "", ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExists
	}

	return val, nil
}

func (s *Store) Set(key, val string) error {
	if key == "" {
		return ErrEmptyKey
	}

	_, ex := s.data[key]
	// check for update else throw error
	if s.maxSize > 0 && s.Len() >= s.maxSize && !ex {
		return fmt.Errorf("Set(%q): %w", key, ErrorStoreFull)
	}

	s.data[key] = val
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
	s.maxSize -= 1
}

func (s *Store) Len() int {
	return len(s.data)
}
