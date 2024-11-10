package main

import "fmt"

func arrayLooping() {
	s := []int{100, 101, 102}

	for v := range s {
		fmt.Println(v) // index
	}

	fmt.Println("----------")
	for _, v := range s {
		fmt.Println(v) // element
	}

	fmt.Println("----------")
	for i, v := range s {
		fmt.Println(i, "=>", v) // index & element
	}

	fmt.Println("----------")
	for i := range s {
		fmt.Println(i) // index
	}
}

/*
	- In for v := range values { the v is the index of the element in the slice.
	- In for _, v := range values { the v is the actual element value.
	- In for i, v := range values { the i is the index and the v is the element.
	- In for i, _ := range values { the i is the index of the element in the slice.
*/
