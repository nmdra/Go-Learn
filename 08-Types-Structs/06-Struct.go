package main

import "fmt"

type Person struct {
	Name     string
	Greet    func(string, string) string
	Farewell func(string) string
}

func main() {
	test := &Person{
		Name: "Test",
		Greet: func(greeting, name string) string {
			return greeting + ", " + name
		},
		Farewell: func(name string) string {
			return "Goodbye, " + name
		},
	}

	fmt.Println(test.Greet("Hi", test.Name))
	fmt.Println(test.Farewell(test.Name))
}
