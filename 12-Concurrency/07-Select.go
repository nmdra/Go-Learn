package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}
