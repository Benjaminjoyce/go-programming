package main

import (
	"fmt"
)

var x int

//int is a whole number
var y float64

//float64 is a number with decimal places
func main() {
	x = 42
	y = 42.3453

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

//byte is an alias for uint8
//rune is an alias for int32
