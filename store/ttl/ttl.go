package ttl

import (
	"fmt"
	"sort"
	"time"

	"github.com/JagdeepSingh13/store"
)

type TtlStore struct {
	data map[string]ttlEntry
	ttl  time.Duration
}

type ttlEntry struct {
	value     string
	expiresAt time.Time
}

func NewTtlStore(ttl time.Duration) *TtlStore {
	return &TtlStore{
		data: make(map[string]ttlEntry),
		ttl:  ttl,
	}
}

func (t *TtlStore) Get(key string) (string, error) {
	if key == "" {
		return "", store.ErrEmptyKey
	}

	entry, ok := t.data[key]
	if !ok || time.Now().After(entry.expiresAt) {
		delete(t.data, key)

		return "", fmt.Errorf("key %s does not exist", key)
	}

	return entry.value, nil
}

func (t *TtlStore) Set(key, val string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	entry := ttlEntry{
		value:     val,
		expiresAt: time.Now().Add(t.ttl),
	}

	t.data[key] = entry
	return nil
}

func (t *TtlStore) Keys() []string {
	keys := make([]string, 0, len(t.data))

	// make slice of keys, sort the keys alb.
	for key := range t.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func (t *TtlStore) Delete(key string) {
	delete(t.data, key)
}

func (t *TtlStore) Len() int {
	return len(t.data)
}
