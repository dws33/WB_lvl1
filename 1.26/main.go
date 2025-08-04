package main

import (
	"fmt"
	"strings"
)

func hasAllUniqueChars(s string) bool {
	seen := make(map[rune]bool)

	s = strings.ToLower(s)

	for _, r := range []rune(s) {
		if seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}

func main() {
	tests := []string{
		"abcd",
		"abCdefAaf",
	}

	for _, str := range tests {
		fmt.Printf("Строка: %-11s Уникальность: %v\n", str, hasAllUniqueChars(str))
	}
}
