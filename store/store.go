package store

import "errors"

var ErrKeyDoesNotExists = errors.New("key does not exists")
var ErrEmptyKey = errors.New("key is mandatory")
var ErrorStoreFull = errors.New("store is full")

type Storer interface {
}
