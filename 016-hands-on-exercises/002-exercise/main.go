// Using the following operators, write expressions and assign their values to variables:
// == , <= , >= , != , > , <
// Now print each of the variables.

package main

import (
	"fmt"
)

func main() {

	a := (42 == 42)
	b := (42 <= 41)
	c := (42 >= 41)
	d := (a != true)
	e := (32 > 31)
	f := (32 < 31)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

// true
// false
// true
// false
// true
// false
