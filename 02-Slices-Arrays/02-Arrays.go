package main

import (
	"fmt"
	"unsafe"
)

func arraysPointer() {
	a := [3]int{99, 100, 101}

	p := unsafe.Pointer(&a[0])

	a1 := unsafe.Pointer(uintptr(p) + 8)
	a2 := unsafe.Pointer(uintptr(p) + 16)

	fmt.Println(*(*int)(p))
	fmt.Println(*(*int)(a1))
	fmt.Println(*(*int)(a2))
}

// Output:
// 99
// 100
// 101
