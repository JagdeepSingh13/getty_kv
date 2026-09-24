package main

import (
	"encoding/base64"
	"fmt"

	"github.com/JagdeepSingh13/store"
	"github.com/JagdeepSingh13/store/kv"
)

func main() {
	cmds := []Command{
		{Op: "SET", Key: "env", Value: "prod"},
		{Op: "SET", Key: "version", Value: "0.0.1"},
		{Op: "SET", Key: "debug", Value: "true"},
		{Op: "GET", Key: "env"},
		{Op: "SET", Key: "region", Value: "us-east-1"},
		{Op: "GET", Key: "version"},
	}

	s := kv.NewStore(0)
	RestoreOnBoot(s, cmds)

	fmt.Println("restore, store size: ", s.Len())
	fmt.Println("keys: ", s.Keys())

	fmt.Println("hello getty")
}

func PopulateDefaults(s store.Storer) error {
	def := map[string]string{
		"env":     "dev",
		"version": "0.0.1",
	}

	for k, v := range def {
		if err := s.Set(k, v); err != nil {
			return fmt.Errorf("Populate Defauls: %w", err)
		}
	}

	return nil
}

func CreateStore() store.Storer {
	plain := kv.NewStore(3)
	logger := NewLoggingMiddleware(plain)

	return logger
}

// need to make Store Interface so that both Store & TtlStore can use
// enc. fn. at same time, Polymorphism

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", nil
	}

	return store.Get(key)
}
