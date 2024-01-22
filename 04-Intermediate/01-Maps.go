package main

import (
	"fmt"
	"maps"
	"reflect"
)

func main() {

	// Create empty Map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// Map Lookup
	// @see :- https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go

	// To test for presence in the map without worrying about the actual value,
	// you can use the blank identifier (_) in place of the usual variable for the value.

	_, prs := m["k2"]
	fmt.Printf("prs: %v (Type: %v)\n", prs, reflect.TypeOf(prs))

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
