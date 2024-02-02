package main

import "fmt"

func getLength(l int) int {
	return l
}

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 { // <--
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("----")

	length := getLength(10)

	if length > 0 {
		fmt.Printf("Valid Length \n")
	}

	// Compact If

	if length := getLength(10); length <= 0 {
		fmt.Printf("Invalid or Zero Length\n")
	} else {
		fmt.Printf("Valid Length")
	}
}

// There is no ternary if in Go,
// so you’ll need to use a full if statement even for basic conditions.
