package main

import "fmt"

type MyInt int
type YourInt int

func PrintInt[T ~int](t T) {
	fmt.Println(t)
}

func main() {
	var a int = 5
	var b MyInt = 10
	var c YourInt = 15

	PrintInt(a)
	PrintInt(b)
	PrintInt(c)
}
