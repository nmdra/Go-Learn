package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Integer types
	var intVar int = 42

	var int8Var int8 = 127
	// var int16Var int16 = 32767
	// var int32Var int32 = 2147483647
	// var int64Var int64 = 9223372036854775807

	// Floating-point types
	var float32Var float32 = 3.14
	// var float64Var float64 = 3.141592653589793

	// Boolean type
	var boolVar bool = true

	// String type
	var stringVar string = "Hello, Go!"

	// Character type (not a separate type in Go, just for illustration)
	var charVar rune = 'A'

	// Complex number types
	var complex64Var complex64 = 1 + 2i
	// var complex128Var complex128 = 3 + 4i

	// Array and Slice types
	var arrayVar [3]int = [3]int{1, 2, 3}
	var sliceVar []int = []int{4, 5, 6}

	// Map type
	var mapVar map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

	// Pointer type
	var pointerVar *int = &intVar

	// Struct type
	type Person struct {
		Name string
		Age  int
	}

	var structVar Person = Person{Name: "John", Age: 30}

	// Displaying the values and types
	display("intVar", intVar)
	display("int8Var", int8Var)
	display("float32Var", float32Var)
	display("boolVar", boolVar)
	display("stringVar", stringVar)
	display("charVar", charVar)
	display("complex64Var", complex64Var)
	display("arrayVar", arrayVar)
	display("sliceVar", sliceVar)
	display("mapVar", mapVar)
	display("pointerVar", pointerVar)
	display("structVar", structVar)
}

// display function prints the value and type of a variable
func display(name string, value interface{}) {
	fmt.Printf("%s: %v (Type: %v)\n", name, value, reflect.TypeOf(value))
	fmt.Println()
}
