package main

import "fmt"

// A higher-order function that takes a function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// A couple of example functions
func add(x, y int) int {
    return x + y
}

func multiply(x, y int) int {
    return x * y
}

func main() {
    fmt.Println("Addition:", applyOperation(3, 4, add))       // Outputs: 7
    fmt.Println("Multiplication:", applyOperation(3, 4, multiply)) // Outputs: 12
}
