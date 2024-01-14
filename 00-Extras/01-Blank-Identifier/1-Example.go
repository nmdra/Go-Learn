package main

import "fmt"

func getData() (int, string) {
	// This function returns two values, but we're not interested in the first one.
	// We use the blank identifier (_) to discard it.
	return 42, "Hello, Gophers!"
}

func main() {
	// Calling the getData function and using the blank identifier to discard the first returned value.
	_, message := getData()

	// Printing only the second returned value.
	fmt.Println("Message:", message)

	// Calling the getData function and capturing the first returned value in the variable result.
	result, _ := getData()

	// Printing only the first returned value.
	fmt.Println("Result:", result)
}
