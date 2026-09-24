package main

import (
	"log"
	"sync"
	"time"

	"github.com/JagdeepSingh13/store"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) {
	switch c.Op {
	case "SET":
		log.Printf("SET key: %s", c.Key)
		time.Sleep(time.Second * 1)
		s.Set(c.Key, c.Value)
	case "GET":
		s.Get(c.Key)
	}
}

// goroutines are virtual threads
// but just adding go abruptly ends even when dispatch() is executing
// as main doesn't wait, so we use WaitGroups

func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup

	for _, c := range cmds {
		wg.Add(1)

		go func() {
			defer wg.Done()
			dispatch(s, c)
		}()
	}

	wg.Wait()
}
