package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 1, 3, 8, 2}
	fmt.Printf("%#v", Intersect(a, b))
}

func Intersect[T comparable](a []T, b []T) []T {
	mp := make(map[T]struct{})
	c := []T{}
	for _, val := range a {
		for _, v := range b {
			if v == val {
				mp[v] = struct{}{}
			}
		}
	}
	for key, _ := range mp {
		c = append(c, key)
	}
	return c
}
