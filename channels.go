package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// channels -> sending data b/w go routines
// chan<- send only channel, <-chan read only

// channels are un-buffered, when full they block, we read before sending another message

func cmain() {
	results := make(chan int, 10)

	go worker(results)
	go worker(results)

	fmt.Println(<-results)
	fmt.Println(<-results)

	fmt.Println("hello")
}

func worker(results chan<- int) {
	time.Sleep(time.Second)

	results <- rand.IntN(100)
}
