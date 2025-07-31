package main

import "fmt"

func main() {
	a := 3
	b := 7
	fmt.Printf("a = %v  b = %v", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	//a = a + b
	//b = a - b
	//a = a - b
	fmt.Printf("\na = %v  b = %v", a, b)
}
