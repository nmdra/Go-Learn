package main

import "fmt"

// Function to modify a value using a pointer
func update(val *int) {
    *val = 100
}

func main() {
    x := 50
    update(&x)
    fmt.Println("Updated value of x:", x) // 100
}
