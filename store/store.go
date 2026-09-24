package store

import "errors"

var ErrKeyDoesNotExists = errors.New("key does not exists")
var ErrEmptyKey = errors.New("key is mandatory")
var ErrorStoreFull = errors.New("store is full")

// need to impl all these methods to use kv, ttl, logger, metrics with Storer
type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Keys() []string
	Delete(key string)
	Len() int
}
