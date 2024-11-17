package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42
	var y int = 50
	var p *int = &x // p now holds the address of x

	fmt.Println("p value: ", reflect.ValueOf(p), "p type: ", reflect.TypeOf(p))

	fmt.Println("Value of x:", x, "Address of x:", &x) // 42
	fmt.Println("Value of p:", p)                      // (same memory address)
	fmt.Println("p:", *p)                              // dereferencing p

	*p = 21 // changing the value at the memory address p points to

	fmt.Println("Value of x after change:", x) // 21

	p = &y

	fmt.Println("p:", *p) // dereferencing p
}
