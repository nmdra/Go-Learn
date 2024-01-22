package main

import "fmt"

// Normal Function
func say_hello(msg string) {
	fmt.Println(msg)
}

func return_anon_func() func(string) {
	// return an anonymous function
	return func(msg string) {
		fmt.Print(msg)
	}
}

func main() {
	say_hello("Hello")

	// anonymous function declared:
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymous")

	print_func := return_anon_func()
	print_func("Hello returned from Anonymous")
}
