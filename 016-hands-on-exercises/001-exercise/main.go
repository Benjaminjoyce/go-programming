//Write a program that prints a number in decimal, binary, and hex
package main

import (
	"fmt"
)

func main() {

	//print a number as a DECIMAL(%d), BINARY(%b) and a HEXI-DECIMAL(#x)
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

}

//42	101010	0x2a
