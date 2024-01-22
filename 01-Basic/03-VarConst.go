package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println()

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.
	// This syntax is only available inside functions.

	f := "apple"
	fmt.Println(f)

	fmt.Println()

	fmt.Println(s)

	const n = 500000000

	const m = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(m))

	fmt.Println(math.Sin(n))
}
