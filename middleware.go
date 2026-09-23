package main

import (
	"log"
	"os"

	"github.com/JagdeepSingh13/store"
)

// so this receives a store
// also we can use this as Stores (since inf. methods are impl in this)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "[log] ", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	l.logger.Printf("GET %q", key)
	val, err := l.inner.Get(key)
	if err != nil {
		l.logger.Printf("GET %q -> miss (%v)", key, err)
	} else {
		l.logger.Printf("GET %q -> hit", key)
	}

	return val, err
}

func (l *LoggingMiddleware) Set(key, val string) error {
	l.logger.Printf("SET %q", key)

	return l.inner.Set(key, val)
}

func (l *LoggingMiddleware) Keys() []string {
	got := l.inner.Keys()
	l.logger.Printf("KEYS %v", got)

	return got
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("DELETE %q", key)

	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Len() int {
	got := l.inner.Len()
	l.logger.Printf("LEN %d", got)

	return got
}
