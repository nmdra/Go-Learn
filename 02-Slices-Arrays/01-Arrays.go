package main

import (
	"fmt"
	"sort"
)

func arraysGo() {

	// Declaring an array and initializing it with values
	var numbers = [...]int{1, 2, 3, 4, 5} // Size calculated automatically

	fmt.Println("numbers:", numbers)

	var array []int
	fmt.Println("array:", array)

	// Sorting Array
	arr := []int{5, 3, 1, 4, 2}
	sort.Ints(arr)

	for i := range arr {
		println(i, &arr[i])
	}

	// byte array
	bytearr := [5]byte{0, 1, 2, 3, 4}
	println("bytearr", &bytearr)

	for i := range bytearr {
		println(i, &bytearr[i])
	}

	target := 3
	found := false
	for _, value := range arr {
		if value == target {
			found = true
			break
		}
	}

	fmt.Println(found)

	fmt.Println()

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	fmt.Println()

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*
	- An array in Go is a fixed-size data structure that stores elements of the same type.
	- You need to specify the size of the array when you declare it.
	- Once an array is defined, its size cannot be changed.
	- Arrays are often used when you know the exact number of elements you need in advance
*/
