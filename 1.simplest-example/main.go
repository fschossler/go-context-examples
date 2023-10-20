package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 2 seconds, if you change to 4 seconds the operation is completed.
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // Ensure the context is canceled when done

	// Simulate a long-running operation
	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Operation completed")
		case <-ctx.Done():
			fmt.Println("Operation canceled due to timeout")
		}
	}()

	// Wait for a few seconds to see the context in action
	time.Sleep(4 * time.Second)
}
