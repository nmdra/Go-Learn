package main

import "fmt"

func main() {
    s := "Hello, නිමෙන්ද්‍ර 😎"
    
    // Convert string to slice of runes
    runes := []rune(s)

    // Print each rune as a character and Unicode code point
    for i, r := range runes {
        fmt.Printf("Character %d: %c, Unicode: %U\n", i, r, r)
    }
}
