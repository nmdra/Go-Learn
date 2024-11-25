package main

import (
	"context"
	"fmt"
)

func main() {
	// Create a context with a key-value pair
	ctx := context.WithValue(context.Background(), "userID", 42)

	// Pass the context to a function
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// Retrieve the value from the context
	userID := ctx.Value("userID")
	if userID == nil {
		fmt.Println("No user ID found in context")
		return
	}
	fmt.Println("Processing request for user ID:", userID)
}
