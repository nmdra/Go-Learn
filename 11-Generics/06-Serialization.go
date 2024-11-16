package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Create an instance of the struct
	person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Serialize the struct to JSON
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error serializing:", err)
		return
	}

	fmt.Println("Serialized JSON:", string(data))
}
