package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 1

	fmt.Println(m)
	fmt.Println(len(m))

	delete(m, "a")
	fmt.Println(m)

	m2 := map[string]int {"c":2, "d":3}
	fmt.Println(m2)

	val, isPresent := m["b"]
	fmt.Println(val, isPresent)

	_, isAPresent := m["a"]
	fmt.Println(isAPresent)
}