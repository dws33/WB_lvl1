package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("Меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a Action) Run() {
	fmt.Printf("%s сейчас бежит!\n", a.Name)
}

func main() {
	a := Action{
		Human: Human{
			Name: "Иван",
			Age:  30,
		},
	}

	a.Speak()
	a.Run()
}
