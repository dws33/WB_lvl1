package main

import (
	"fmt"
)

func setBit(n int64, i uint, bitValue int) int64 {
	if bitValue == 1 {
		n = n | (1 << i) //установка i-того бита в 1
	} else {
		n = n &^ (1 << i) //установка i-того бита в 0
	}
	return n
}

func main() {
	var number int64 = 5
	var bitIndex uint = 1
	var newValue int = 1
	result := setBit(number, bitIndex, newValue)
	fmt.Printf("Было:  %09b\nСтало: %09b", number, result)
}
