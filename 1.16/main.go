package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]

	var less []int
	var greater []int

	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	arr := []int{9, 2, 7, 4, 5, 1, 8, 3, 6}
	fmt.Println("Sorted:", quickSort(arr))
}
