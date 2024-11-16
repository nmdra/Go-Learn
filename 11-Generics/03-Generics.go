package main

import (
	"fmt"
)

type MyConstraint interface {
	int | int8 | int16 | int32 | int64
}

// Example of a generic function using
// the MyConstraint type parameter
func MyFunc[T MyConstraint](input T) {
	fmt.Printf("Input type: %T, Input value: %v\n", input, input)
}

func main() {
	// Valid because type int8 is in the type set
	// (int | int8 | int16 | int32 | int64)
	MyFunc[int8](1)

	// Invalid because type string is not in the type set
	//   MyFunc[string]("hello")
}
