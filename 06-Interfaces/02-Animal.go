package main

import "fmt"

type Animal interface {
	Speak() string
	// Eats() string // 😮 must implement
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

// func (d Dog) Eats() string {
//     return "Dog food"
// }

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type GoProgrammer struct {
}

func (j GoProgrammer) Speak() string {
	return "Interfaces"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, GoProgrammer{}} // Slice
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
