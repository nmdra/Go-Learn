package main

import "fmt"

// type empty interface {}

// func printAnything(x empty) {
//     fmt.Println(x)
// }

func printAnything(x... interface{}) {
    fmt.Println(x)
}



func main() {
    printAnything(42)        // Passes an int
    printAnything("Hello")   // Passes a string
    printAnything(true, 11, "Hii")      // Passes a bool

	mysteryBox := interface{}("String")
	describeValue(mysteryBox) 

	retrieveValue, ok := mysteryBox.(int)
	if ok {
        fmt.Printf("Retrieved value is an int: %v\n", retrieveValue)
    } else {
        fmt.Println("Retrieved value is not an int")
    }
}

func describeValue(v interface{}) {
	switch v.(type) {
    case int:
        fmt.Printf("I'm an int and it's value is %v\n", v)
    case string:
        fmt.Printf("I'm a string and it's value is %v\n", v)
    case bool:
        fmt.Printf("I'm a bool and it's value is %v\n", v)
    default:
        fmt.Printf("I don't know the type of this value\n")
    }
}