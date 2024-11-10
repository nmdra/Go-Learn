package main

import "fmt"

func sliceAndArray() {
	array := [6]int{0, 1, 2, 3, 4, 5}

	slice := array[1:3]
	fmt.Println(slice, len(slice), cap(slice))

	slice[0] = 100

    fmt.Println(array) // Output: [0 100 2 3 4 5]

	fmt.Println(&array[1], ",", &slice[0]) // same

	// slice = append(slice, 6, 7, 8, 9)

	slice = append(slice, 6, 7, 8)

	fmt.Println(slice, len(slice), cap(slice)) // Output: [1 2 6 7] 4 8
	fmt.Println(array) // Output: [0 100 2 6 7 5]
}