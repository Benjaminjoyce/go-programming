package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	d
	e
	f
)

const (
	g = iota
	h
	i
	j
	//starting a new const() will reset ioto from incrementing
)

const (
	red = iota
	blue
	green
	purple
	cake
	grass
)

const nwln = "This is the start of a new set of constants"

func main() {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(nwln)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(nwln)
	fmt.Println(red)
	fmt.Println(blue)
	fmt.Println(green)
	fmt.Println(purple)
	fmt.Println(cake)
	fmt.Println(grass)
}

// int
// 0
// 1
// 2
// 3
// 4
// 5
// This is the start of a new set of constants
// 0
// 1
// 2
// 3
// This is the start of a new set of constants
// 0
// 1
// 2
// 3
// 4
// 5
