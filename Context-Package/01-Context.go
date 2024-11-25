package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Cancel the context to free up resources

	// Launch a goroutine
	go func(ctx context.Context) {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Task completed")
		case <-ctx.Done():
			fmt.Println("Task canceled:", ctx.Err())
		}
	}(ctx)

	// Wait for the goroutine to finish or timeout
	time.Sleep(3 * time.Second)
}
