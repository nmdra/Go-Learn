package main

import "fmt"

// Define a struct type
type Person struct {
    Name string
    Age  int
}

// Define a method for the Person type
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s, and I am %d years old.\n", p.Name, p.Age)
}

func (p *Person) updateAge(age int) {
	p.Age = age
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    person.Greet()  // Calling the Greet method on a Person instance

	person.updateAge(32)

	person.Greet()
}
