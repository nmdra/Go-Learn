package main

import "fmt"

func main() {

	// create an empty interface
	var a interface{}
	a = 12

	// type assertion
	interfaceValue, _ := a.(string)

	fmt.Println("Interface Value:", interfaceValue)
	// fmt.Println("Boolean Value:", booleanValue)
}
