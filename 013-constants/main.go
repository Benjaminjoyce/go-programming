package main

import (
	"fmt"
)

const (
	a = 42
	b = 42.78
	c = "Harry Potter"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

// 42
// 42.78
// Harry Potter
// int
// float64
// string
