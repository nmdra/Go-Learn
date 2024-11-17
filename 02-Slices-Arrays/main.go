package main

import (
	"fmt"
	"os"
)

func main() {
	// Displaying a menu for the user
	fmt.Println("I want to be a Gopher!")
	fmt.Println("Choose a function to run:")
	fmt.Println("1. Arrays")
	fmt.Println("2. Array with pointers")
	fmt.Println("3. Slices")
	fmt.Println("4. Looping through Arrays")
	fmt.Println("5. Slice and Array")
	fmt.Println("6. How Slice Grow")

	// Get user input
	var choice int
	fmt.Print("Enter the number of the function you want to run: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input, exiting.")
		os.Exit(1)
	}

	fmt.Println("================================================")

	// Switch case to run the selected function
	switch choice {
	case 1:
		arraysGo()
	case 2:
		arraysPointer()
	case 3:
		slicesGo()
	case 4:
		arrayLooping()
	case 5:
		sliceAndArray()
	case 6:
		sliceGrow()
	default:
		fmt.Println("Invalid choice, exiting.")
		os.Exit(1)
	}
}
