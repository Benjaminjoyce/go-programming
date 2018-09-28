//Create a variable of type string using a raw string literal. Print it.

package main

import (
	"fmt"
)

var x string = "this is a raw string litteral"

func main() {

	x := `
"this is a raw string litteral"
we can put anything in here

including line breaks
( 8

`
	fmt.Println(x)
}

// "this is a raw string litteral"
// we can put anything in here

// including line breaks
// ( 8
