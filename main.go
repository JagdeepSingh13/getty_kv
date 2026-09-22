package main

import (
	"errors"
	"fmt"
)

func main() {
	s := NewStore()

	s.Set("a", "45")
	s.Delete("a")

	val, err := s.Get("b")
	if err != nil {
		if errors.Is(err, ErrKeyDoesNotExists) {
			// do smtg like -> 400 http
			fmt.Println("key not exists error")
			return
		}
		// 500
		fmt.Println(err)
	}
	fmt.Println(val)

	fmt.Println("hello getty")
}
