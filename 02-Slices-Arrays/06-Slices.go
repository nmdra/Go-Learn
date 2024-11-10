package main

import (
    "fmt"
    "unsafe"
)

func sliceGrow() {
	array := [6]int{0, 1, 2, 3, 4, 5}

	fmt.Println("array:", &array[1])

	slice := array[1:3:4]
	fmt.Println("slice:", slice)
	fmt.Println("slice:", &slice[0])

	slice = append(slice, 6)
	fmt.Println("slice after appending 6:", slice, unsafe.SliceData(slice))
	fmt.Println("array now:", array)
	fmt.Println("slice now:", &slice[0])

	slice = append(slice, 7)
	fmt.Println("slice after appending 7:", slice, unsafe.SliceData(slice))
	fmt.Println("slice:", &slice[0])

	fmt.Println("array now:", array)
}

// Output:
// slice: [1 2]
// slice after appending 6: [1 2 6] 0x14000128038
// array now: [0 1 2 6 4 5]
// slice after append 7: [1 2 6 7] 0x14000128090
// array now: [0 1 2 6 4 5]