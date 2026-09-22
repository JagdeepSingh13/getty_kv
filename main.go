package main

import (
	"fmt"
)

func main() {
	s := NewStore()

	s.Set("a", "45")
	s.Set("b", "50")

	fmt.Println(s.Get("a"))
}
