package kv

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/JagdeepSingh13/store"
)

// maps in go are not concurrent
type Store struct {
	mu      sync.Mutex
	data    map[string]string
	maxSize int
}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize, // 0 -> unlimited
	}
}

// need to make a copy of the store, but since data(hmap) is a ptr
// copy store modifies the org. store, so used loop
func (s *Store) Clone() *Store {
	cp := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for k, v := range s.data {
		cp.data[k] = v
	}

	return cp
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
	s.mu.Lock()
	defer s.mu.Unlock()

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

	s.mu.Lock()
	defer s.mu.Unlock()

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

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ex := s.data[key]
	// check for update key else throw error since full cap.
	// we changed from s.Len() to len() as its also has mu so a deadlock condn
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !ex {
		return fmt.Errorf("Set(%q): %w", key, store.ErrorStoreFull)
	}

	s.data[key] = val
	return nil
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
	s.maxSize -= 1
}

func (s *Store) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.data)
}

func (s *Store) Incr(key string) (int, error) {
	if key == "" {
		return 0, store.ErrEmptyKey
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	cur := 0
	if raw, ok := s.data[key]; ok {
		v, err := strconv.Atoi(raw)
		if err != nil {
			return 0, fmt.Errorf("INCR %q: %w", key, err)
		}
		cur = v
	}

	cur++
	time.Sleep(time.Microsecond)

	s.data[key] = strconv.Itoa(cur)

	return cur, nil
}
