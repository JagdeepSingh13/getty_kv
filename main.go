package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/JagdeepSingh13/store"
	"github.com/JagdeepSingh13/store/kv"
	"github.com/JagdeepSingh13/store/ttl"
)

func main() {
	plain := kv.NewStore(3)
	logger := NewLoggingMiddleware(plain)
	metrics := NewMetricsMiddleware(logger)
	// s := CreateStore()

	_, err := SetKeyWithEncryption(metrics, "d", "60")
	if err != nil {
		fmt.Println(err)
		return
	}

	metrics.Report()

	ttlStore := ttl.NewTtlStore(time.Second * 2)
	encc, err := SetKeyWithEncryption(ttlStore, "e", "ttl entry")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encc)

	fmt.Println("hello getty")
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
