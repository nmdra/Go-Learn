package main

import (
	"fmt"
	"testing"
	"time"
)

// Define a larger struct with multiple fields
type Person struct {
	Name       string
	Age        int
	Occupation string
	Address    string
	Phone      string
	Email      string
	Interests  []string
	Skills     [5]string
}

// Function that takes a large struct by value (copy)
func modifyPersonCopy(p Person) {
	p.Name = "Modified"
	p.Age = 50
	p.Occupation = "Engineer"
}

// Function that takes a pointer to a large struct
func modifyPersonPointer(p *Person) {
	p.Name = "Modified"
	p.Age = 50
	p.Occupation = "Engineer"
}

// Benchmark function to measure performance with copy
func BenchmarkModifyPersonCopy(b *testing.B) {
	person := Person{
		Name:       "John Doe",
		Age:        30,
		Occupation: "Teacher",
		Address:    "123 Main St",
		Phone:      "555-5555",
		Email:      "johndoe@example.com",
		Interests:  []string{"Reading", "Traveling", "Music"},
		Skills:     [5]string{"Writing", "Teaching", "Public Speaking", "Management", "Analysis"},
	}

	for i := 0; i < b.N; i++ {
		modifyPersonCopy(person)
	}
}

// Benchmark function to measure performance with pointer
func BenchmarkModifyPersonPointer(b *testing.B) {
	person := Person{
		Name:       "John Doe",
		Age:        30,
		Occupation: "Teacher",
		Address:    "123 Main St",
		Phone:      "555-5555",
		Email:      "johndoe@example.com",
		Interests:  []string{"Reading", "Traveling", "Music"},
		Skills:     [5]string{"Writing", "Teaching", "Public Speaking", "Management", "Analysis"},
	}

	for i := 0; i < b.N; i++ {
		modifyPersonPointer(&person)
	}
}

func main() {
	// Run benchmarks manually
	person := Person{
		Name:       "John Doe",
		Age:        30,
		Occupation: "Teacher",
		Address:    "123 Main St",
		Phone:      "555-5555",
		Email:      "johndoe@example.com",
		Interests:  []string{"Reading", "Traveling", "Music"},
		Skills:     [5]string{"Writing", "Teaching", "Public Speaking", "Management", "Analysis"},
	}

	start := time.Now()
	for i := 0; i < 100000000000; i++ {
		modifyPersonCopy(person)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken by copy: %v\n", elapsed)

	start = time.Now()
	for i := 0; i < 100000000000; i++ {
		modifyPersonPointer(&person)
	}
	elapsed = time.Since(start)
	fmt.Printf("Time taken by pointer: %v\n", elapsed)
}

// run: go test -bench=.
