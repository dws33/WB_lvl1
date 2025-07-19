package main

import "fmt"

func numGen(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func numSquarer(nums <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range nums {
			out <- n * n
		}
	}()

	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for n := range numSquarer(numGen(nums)) {
		fmt.Println(n)
	}

}
