package main

import "fmt"

const (
	A = iota // A == 0
	B        // B == 1
	C        // C == 2
)

func main() {
	fmt.Println(A, B, C) // Output: 0 1 2
}
