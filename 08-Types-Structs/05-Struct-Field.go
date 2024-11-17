package main

import "fmt"

type Person struct {
	Name     string
	Greet    func(string) string
	Farewell func() string
}

func main() {
	person := Person{
		Name: "C",
	}
	// Assign the greet and farewell functions after the person is defined
	person.Greet = func(greeting string) string {
		return greeting + ", " + person.Name
	}
	person.Farewell = func() string {
		return "Goodbye, " + person.Name
	}
	// Call the function fields
	fmt.Println(person.Greet("Hello"))
	fmt.Println(person.Farewell())
}
