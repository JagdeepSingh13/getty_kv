package kv

import (
	"fmt"
	"sort"

	"github.com/JagdeepSingh13/store"
)

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
		return "", store.ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", store.ErrKeyDoesNotExists
	}

	return val, nil
}

func (s *Store) Set(key, val string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	_, ex := s.data[key]
	// check for update key else throw error since full cap.
	if s.maxSize > 0 && s.Len() >= s.maxSize && !ex {
		return fmt.Errorf("Set(%q): %w", key, store.ErrorStoreFull)
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
