package main

import "fmt"

func main() {
    person := struct {
        Name string
        Age  int
    }{
        Name: "Alice",
        Age:  25,
    }
    fmt.Println(person) // Output: {Alice 25}
}