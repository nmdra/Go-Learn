package main

import "fmt"

func genericAddNums[N int | float64](n1, n2 N) N {
	return n1 + n2
}

func main() {
	fmt.Println(genericAddNums(1, 2))
	fmt.Println(genericAddNums(1.1, 2.2))
}
