package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second) // Simulate work

	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Set the number of logical CPUs to use
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Number of CPUs:", runtime.NumCPU())

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed")
}
