package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Generic function to find the minimum value
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Min of 3 and 5:", Min(3, 5))            // Output: 3
	fmt.Println("Min of 1.1 and 2.2:", Min(1.1, 2.2))    // Output: 1.1
	fmt.Printf("Min of 'a' and 'b':%c\n", Min('a', 'b')) // Output: a
}
