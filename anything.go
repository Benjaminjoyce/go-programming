package main

import (
	"fmt"
)

func main() {
	fmt.Println("hellow everyone, this is the go programming course")
	foo()
	fmt.Println("yeah this is just some more stuff")
	foo()
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo!")
}

func bar() {
	fmt.Println("end of func main")
}
