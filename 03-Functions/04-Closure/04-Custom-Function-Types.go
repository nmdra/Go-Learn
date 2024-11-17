package main

import "fmt"

type MathOperation func(int, int) int

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }

func main() {
	var op MathOperation

	op = add
	fmt.Println("Add:", op(5, 3)) // Output: Add: 8

	op = subtract
	fmt.Println("Subtract:", op(5, 3)) // Output: Subtract: 2
}
