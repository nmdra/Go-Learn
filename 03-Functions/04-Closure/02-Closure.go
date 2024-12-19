package main

import "fmt"

// A function that returns a closure
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter1 := createCounter()
    counter2 := createCounter()

    fmt.Println(counter1()) // Outputs: 1
    fmt.Println(counter1()) // Outputs: 2

    fmt.Println(counter2()) // Outputs: 1
    fmt.Println(counter2()) // Outputs: 2
}
