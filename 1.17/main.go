package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 4, 6, 7, 10, 12, 16, 25}
	target := 12

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d находится по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
