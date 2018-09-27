package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	//%b	base 2
	fmt.Printf("%x\n", y)
	//%x	base 16, with lower-case letters for a-f
	fmt.Printf("%#x\n", y)
	//%#x	base 16, with lower-case letters for a-f with prepended 0x
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)
	// for each fm.Printf arguement a value needs to be passed.

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)

	fmt.Printf("%T\n", s)
	fmt.Println("s:", s)
	fmt.Printf("%v", y)

}
