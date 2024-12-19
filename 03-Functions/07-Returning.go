package main

import "fmt"

// A higher-order function that returns another function
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)

    fmt.Println("Double 5:", double(5)) // Outputs: 10
    fmt.Println("Triple 5:", triple(5)) // Outputs: 15
}
