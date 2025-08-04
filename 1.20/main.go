package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseRunes(runes []rune, i, j int) {
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
}

func reverseWordsInPlace(s string) string {
	runes := []rune(s)

	reverseRunes(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		result := reverseWordsInPlace(strings.TrimSpace(input))
		fmt.Println("Результат:", result)
	}
}
