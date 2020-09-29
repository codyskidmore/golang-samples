package main

import (
	"fmt"
)

func main() {
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	fmt.Println(a << 3)
	fmt.Println(a >> 3)

	n := 3.14
	fmt.Println(n)

	n = 13.7e72
	fmt.Println(n)

	n = 2.1E14
	fmt.Printf("%v, %T\n", n, n)

	var c complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", c, c)
}