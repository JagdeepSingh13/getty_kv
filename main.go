package main

import (
	"errors"
	"fmt"
)

func main() {
	// need to pass the max. keys we can store
	// also can do Set() again on same key even if full cap., like a update
	s := NewStore(2)

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
