package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Print("Введите строку: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		reversed := reverseString(input)
		fmt.Println("Перевёрнутая строка:", reversed)
	}
}
