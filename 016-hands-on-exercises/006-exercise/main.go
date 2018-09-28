// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import (
	"fmt"
)

const (
	_ = iota
	a = (2018 + iota)
	b = (2018 + iota)
	c = (2018 + iota)
	d = (2018 + iota)
)

func main() {

	fmt.Println(a, b, c, d)
}

//2019 2020 2021 2022
