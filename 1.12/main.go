package main

import "fmt"

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	mp := make(map[string]struct{})
	for _, val := range sl {
		mp[val] = struct{}{}
	}
	res := []string{}
	for key, _ := range mp {
		res = append(res, key)
	}
	fmt.Println(res)
}
