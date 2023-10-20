package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a context with a 3-second timeout, if we change this to 1 millisecond
	// we receive the error 'context deadline exceeded'.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure the context is canceled when done

	// Create an HTTP client with the context
	client := &http.Client{}

	// Create an HTTP request
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Associate the context with the request
	req = req.WithContext(ctx)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful")
	} else {
		fmt.Println("Request failed with status code:", resp.Status)
	}
}
