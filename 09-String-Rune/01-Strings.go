package main

import "fmt"

func main() {
	s := "Hello, 世界"    // A string containing ASCII and Unicode characters
	fmt.Println(s)      // Output: Hello, 世界
	fmt.Println(len(s)) // Output: 13 (in bytes, not characters)
}
