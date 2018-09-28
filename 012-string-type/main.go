package main

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	b := `Hello,
	playground`
	fmt.Println(s)
	fmt.Println(b)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
	g := "H"
	fmt.Println(g)

	gh := []byte(g)
	fmt.Println(gh)

	n := gh[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
}

// Hello, playground
// Hello,
// 	playground
// string
// [72 101 108 108 111 44 32 112 108 97 121 103 114 111 117 110 100]
// []uint8
// U+0048 'H'
// U+0065 'e'
// U+006C 'l'
// U+006C 'l'
// U+006F 'o'
// U+002C ','
// U+0020 ' '
// U+0070 'p'
// U+006C 'l'
// U+0061 'a'
// U+0079 'y'
// U+0067 'g'
// U+0072 'r'
// U+006F 'o'
// U+0075 'u'
// U+006E 'n'
// U+0064 'd'
// at index position 0 we have hex 0x48
// at index position 1 we have hex 0x65
// at index position 2 we have hex 0x6c
// at index position 3 we have hex 0x6c
// at index position 4 we have hex 0x6f
// at index position 5 we have hex 0x2c
// at index position 6 we have hex 0x20
// at index position 7 we have hex 0x70
// at index position 8 we have hex 0x6c
// at index position 9 we have hex 0x61
// at index position 10 we have hex 0x79
// at index position 11 we have hex 0x67
// at index position 12 we have hex 0x72
// at index position 13 we have hex 0x6f
// at index position 14 we have hex 0x75
// at index position 15 we have hex 0x6e
// at index position 16 we have hex 0x64
// H
// [72]
// 72
// uint8
// 1001000
// 0X48
