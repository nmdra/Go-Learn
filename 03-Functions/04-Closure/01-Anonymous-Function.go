package main

import (
	"fmt"
	"reflect"	
)

// Normal Function
func say_hello(msg string) {
	fmt.Println(msg)
}

func return_anon_func() func(string) {
	// return an anonymous function
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {
	say_hello("Hello")

	// anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymous")

	print_func := return_anon_func()

	fmt.Println(reflect.ValueOf(print_func))
	fmt.Println(reflect.TypeOf(print_func))

	print_func("Hello returned from Anonymous")
}
