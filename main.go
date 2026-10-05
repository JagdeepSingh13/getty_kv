package main

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/JagdeepSingh13/store"
	"github.com/JagdeepSingh13/store/kv"
)

func main() {
	store := kv.NewStore(0)
	store.Set("name", "jsingh")

	// using RWMutex
	// so that many go routines can read simultaneously but
	// Lock when write oprn
	const readers = 100
	const readEach = 10

	start := time.Now()
	var wg sync.WaitGroup
	for range readers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range readEach {
				store.Get("name")
			}
		}()
	}
	wg.Wait()

	fmt.Printf("%d readers did %d cons. cuncurr. reads in %s\n", readers, readEach,
		time.Since(start).Round(time.Millisecond))

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
