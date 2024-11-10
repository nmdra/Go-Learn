package main

import (
	"fmt"
	"slices"
)

func slicesGo() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Create a slice with initial length 5 and capacity 10
	numbers := make([]int, 5, 10)
    fmt.Println("numbers:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	// Append elements to the slice
	numbers = append(numbers, 1, 2, 3, 4, 5)
    fmt.Println("numbers:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	numbers = append(numbers, 1)
    fmt.Println("numbers:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	numbers = append(numbers, 6,7,8, 9, 10, 11)
	fmt.Println("numbers:", numbers, "len:", len(numbers), "cap:", cap(numbers))


}

/*
	- A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an underlying array.
	- Slices are created using the make() function or by slicing an existing array, another slice, or a string.
	- Slices can grow or shrink as needed.
	- Slices are more commonly used in Go because they provide greater flexibility when working with collections of data.
*/
