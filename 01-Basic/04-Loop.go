package main

import "fmt"

func loop() {

	// Simple loop using a for loop with a condition
	i := 1
	for i <= 3 {
		fmt.Print(i,",")
		i = i + 1
	}

	// Classic for loop with initialization, condition, and increment statement
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Infinite loop with a break statement
	for {
		fmt.Println("loop")
		break // This breaks out of the infinite loop
	}

	// Loop with continue statement to skip even numbers
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // Skip even numbers, go to the next iteration
		}
		fmt.Println(n)
	}

	fmt.Println()

	// For-each range loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// with blank identifiers
	for _,ch := range "hello" {
		fmt.Printf("%c,",ch)
	}
}
