package main

import (
	"fmt"
	"os"
)

func main() {
	// Displaying a menu for the user
	fmt.Println("I want to be a Gopher!")
	fmt.Println("Choose a function to run:")
	fmt.Println("1. Data Types")
	fmt.Println("2. Variables & Constants")
	fmt.Println("3. Loop")
	fmt.Println("4. If")
	fmt.Println("5. Switch")
	fmt.Println("6. Switch 2")

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
		dataTypes()
	case 2:
		varConst()
	case 3:
		loop()
	case 4:
		ifGo()
	case 5:
		switchGo()
	case 6:
		switchGo2()
	default:
		fmt.Println("Invalid choice, exiting.")
		os.Exit(1)
	}
}
