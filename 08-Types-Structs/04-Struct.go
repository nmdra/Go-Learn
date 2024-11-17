package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	john := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}
	fmt.Println(john.Name) // Output: John Doe
	fmt.Println(john)      // Output: 30
}
