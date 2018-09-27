package main

import "fmt"

var h = 55

//DECLARE there is a VARIABLE with the IDENTIFIER "z"
//and that the VARIABLE  with the IDENTIFIER "z" is of TYPE int
//ASSIGNS the ZERO VALUE of TYPE in to "z"

//"no explicit initialization is provided, the variable or value
// is given a default value. Each element of such a variable or value
//is set to the zero value for its type: false for booleans,
//0 for numeric types, "" for strings, and nil for pointers, functions,
//interfaces, slices, channels, and maps." -https://golang.org/ref/spec#The_zero_value

var z int

func main() {
	//short declaration operator ":="
	//DECLARE a variable and ASSIGN a VALUE(of a certain TYPE)
	x := 42
	fmt.Println(x)
	var y = 43
	fmt.Println(y)
	fmt.Println(h)
	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println("again: ", h)
}
