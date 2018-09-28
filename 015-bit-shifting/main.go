package main

import (
	"fmt"
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a)
	b := a << 4
	fmt.Printf("%d\t\t%b\n", b, b)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}

// 4			100
// 64			1000000
// 1024			10000000000
// 1048576		100000000000000000000
// 1073741824	1000000000000000000000000000000

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func second() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}

// 1024		10000000000
// 1048576		100000000000000000000
// 1073741824	1000000000000000000000000000000
