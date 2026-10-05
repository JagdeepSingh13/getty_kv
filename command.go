package main

import (
	"fmt"
	"sync"

	"github.com/JagdeepSingh13/store"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) error {
	switch c.Op {
	case "SET":
		// log.Printf("SET key: %s", c.Key)
		// time.Sleep(time.Second * 1)
		return s.Set(c.Key, c.Value)
	case "INCR":
		_, err := s.Incr(c.Key)
		return err
	default:
		return fmt.Errorf("unknown operation %q", c.Op)
	}
}

// goroutines are virtual threads
// but just adding go abruptly ends even when dispatch() is executing
// as main doesn't wait, so we use WaitGroups

func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	// bufferedso no bottleneck of waiting
	errs := make(chan error, len(cmds))

	for _, c := range cmds {
		wg.Add(1)

		go func() {
			defer wg.Done()
			errs <- dispatch(s, c)
		}()
	}

	// go func -> so that we can do above and below for loops at same time
	// then close wg and channel
	go func() {
		wg.Wait()
		close(errs)
	}()

	// log the errors we get from a channel
	for err := range errs {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
