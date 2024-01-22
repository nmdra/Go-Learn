package main

import "fmt"

func main() {
	// Your code goes here inside the main function
	var ptr *int
	fmt.Println(ptr) // Outputs: <nil>

	// var n = nil // nil is untyped

	isNil()
}

func isNil() {
	var i *int = nil
	var j *int = nil
	if i == j {
		fmt.Println("equal")
	} else {
		fmt.Println("no")
	}
}

/*
	nil in Go is simply the NULL pointer value of other languages.
	nil is a predeclared identifier representing the zero value for pointers,
	channels, functions, interfaces, maps, and slices. It is a constant with a specific type,
	and it is often used to represent the absence of a value or a zero value for certain types

	@see :- https://go101.org/article/nil.html
*/
