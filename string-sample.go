package main

import (
	"fmt"
)

func main() {
	// strings are immutable.
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))

	// Convert to UTF bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
