package main

import "fmt"

func RemoveAt[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		panic("индекс вне диапазона")
	}
	copy(slice[index:], slice[index+1:])

	var zero T
	slice[len(slice)-1] = zero

	return slice[:len(slice)-1]
}

func main() {
	ints := []int{10, 20, 30, 40}
	ints = RemoveAt(ints, 2)
	fmt.Println("ints:", ints)

	strs := []string{"apple", "banana", "cherry", "date"}
	strs = RemoveAt(strs, 1)
	fmt.Println("strs:", strs)

	type User struct{ Name string }
	users := []*User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Carol"},
	}
	users = RemoveAt(users, 0)
	fmt.Println("users:", users[0].Name)
}
