package main

import (
    "fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 11
	m["k3"] = 15
	m["k3"] = 20

	fmt.Printf("map: %v, len: %d\n", m, len(m))

	v3 := m["k3"]
    fmt.Println("v3:", v3)

	v2, k2 := m["k2"]
    fmt.Println("v2:", v2, "k2:", k2)

	delete(m, "k1")

	fmt.Printf("map: %v, len: %d\n", m, len(m))

	clear(m)
	fmt.Printf("map: %v, len: %d\n", m, len(m))

	_, prs := m["k2"]
    fmt.Println("prs:", prs)

	// ----------------------------------------------------------------

	// Map literals 
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

	n3 := maps.Clone(n2)
	n3["foo"] = 5
    fmt.Println(maps.Equal(n2, n3))



}