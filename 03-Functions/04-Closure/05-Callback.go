package main

import (
	"fmt"
	"strings"
)

type Callback func(string)

func process(data string, cb Callback) {
	fmt.Println("Processing:", data)
	cb(data) // Calling the callback after processing
	fmt.Println("Processed:", data)
}

// Alternative way

// func process(data string, cb func(string)) {
//     fmt.Println("Processing:", data)
//     cb(data)  // Calling the callback after processing
// 	fmt.Println("Processed:", data)
// }

func printUpperCase(data string) {
	fmt.Println("Uppercase:", strings.ToUpper(data))
}

func printLowerCase(data string) {
	fmt.Println("Lowercase:", strings.ToLower(data))
}

func main() {
	process("hello", printUpperCase) // Output: Processing: hello

	process("Hello", printLowerCase)

}
