package main

import (
	"fmt"
	"log"

	"github.com/mikeschinkel/go-jsontest"
	_ "github.com/mikeschinkel/go-jsontest/pipefuncs" // Required for pipe functions
)

func main() {
	// Example JSON response
	body := []byte(`{
		"jsonrpc": "2.0",
		"result": {
			"status": "success",
			"content": [
				{"type": "text", "text": "Hello"},
				{"type": "text", "text": "World"}
			],
			"count": 2
		}
	}`)

	// Define assertions using declarative map syntax
	err := jsontest.TestJSON(body, map[string]any{
		// Simple path assertions
		"jsonrpc":          "2.0",
		"result.status":    "success",
		"result.count":     2,
		"result.content.#": 2, // Array length

		// Type markers
		"result.content.0": jsontest.NotNull{},

		// Pipe functions for validation
		"result.content.0.text|notEmpty()": true,
		"result.content|len()":             2,

		// Array collection with order-insensitive comparison
		"result.content.[].type": jsontest.AnyOrder("text", "text"),
		"result.content.[].text": jsontest.AnyOrder("Hello", "World"),
	})

	if err != nil {
		log.Fatalf("JSON assertions failed: %v", err)
	}

	fmt.Println("All JSON assertions passed!")

	// Example with nested objects
	user := []byte(`{
		"user": {
			"id": 123,
			"name": "Alice",
			"email": "alice@example.com",
			"active": true
		}
	}`)

	err = jsontest.TestJSON(user, map[string]any{
		"user.id":     123,
		"user.name":   "Alice",
		"user.active": true,
		"user.email":  jsontest.NotEmpty{},
	})

	if err != nil {
		log.Fatalf("User assertions failed: %v", err)
	}

	fmt.Println("User JSON assertions passed!")
}
