package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 10
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg := <-ch:
			fmt.Println("Received:", msg)
		default:
			fmt.Println("No channel ready")
		}

		time.Sleep(200 * time.Millisecond)
	}
}
