//Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import (
	"fmt"
)

const (
	a int = 54
	b     = 55
)

func main() {
	fmt.Println(a, "\t", b)
	fmt.Printf("%T\t%T\n", a, b)
}
