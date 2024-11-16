package main

import "fmt"

// Define an enum type using iota
type Day int

const (
    Sunday Day = iota  // 0
    Monday             // 1
    Tuesday            // 2
    Wednesday          // 3
    Thursday           // 4
    Friday             // 5
    Saturday           // 6
)

func main() {
    // Print each day
    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

    // Example usage
    today := Wednesday
    fmt.Println("Today is:", today)
}
