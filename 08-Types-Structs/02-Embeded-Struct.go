package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address           // Embedded struct (no explicit field name)
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.City)      // Accessing embedded field directly
	fmt.Println("Country:", p.Country)
}