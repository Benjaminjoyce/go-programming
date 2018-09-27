package main

import (
	"fmt"
)

var a int

type wizard int

var b wizard

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//a = b does not work, fmt.Printf("%T", a) == int, fmt.Printf("%T", b) == wizard
	//once a TYPE is DECLARED you cannot assign it to a different TYPE

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	//Conversions are expressions of the form T(x) where T is a type and x
	//is an expression that can be converted to type T.

}
