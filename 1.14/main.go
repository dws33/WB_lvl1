package main

import (
	"fmt"
	"reflect"
)

func checkType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("value is int: %v\n", v)
	case string:
		fmt.Printf("value is string: %q\n", v)
	case bool:
		fmt.Printf("value is bool: %v\n", v)
	default:
		// Проверим, является ли value каналом
		//interface{} в рантайме теряет информацию о типе канала (поэтому нельзя обработать через type switch)
		t := reflect.TypeOf(value)
		if t.Kind() == reflect.Chan {
			fmt.Printf("value is channel: %v\n", t)
		} else {
			fmt.Printf("unknown type: %T\n", value)
		}
	}
}

func main() {
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(make(chan int))
}
