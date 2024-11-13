package main

import "fmt"

// Function that takes a pointer to an integer
func printValue(val *int) {
	// Check if the pointer is nil to prevent runtime errors
	if val == nil {
		fmt.Println("Pointer is nil, cannot dereference.")
		return
	}
	// Dereference the pointer to access the integer value
	fmt.Println("Value:", *val)
}

func main() {
	// Declare a nil pointer to an integer
	var ptr *int

	// Try to pass the nil pointer to the function
	printValue(ptr) // Output: Pointer is nil, cannot dereference.
}
