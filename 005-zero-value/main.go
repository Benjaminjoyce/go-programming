package main

import "fmt"

var y string
var z int

func main() {
	//DECLARE a VARIABLE to be of a certain TYPE
	//and then ASSIGN a VALUE of that type to the VARIABLE

	fmt.Println("printing the value of y: ", y, " ending")
	fmt.Printf("%T\n", y)

	y = "Your a wizard Harry"

	fmt.Println("printing the value of y: ", y, " ending")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T", z)
}
