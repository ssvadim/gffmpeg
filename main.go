package main

import (
	"fmt"
	"net/http"
	"os"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the message from query parameters
	message := r.URL.Query().Get("message")
	if message == "" {
		message = "Hello, World!"
	}

	// Fetch the TEST environment variable
	testEnv := os.Getenv("TEST")
	if testEnv == "" {
		testEnv = "Environment variable TEST is not set"
	}

	// Respond to the client
	fmt.Fprintf(w, "Echo: %s\nTEST: %s\n", message, testEnv)
}

func main() {
	http.HandleFunc("/", echoHandler)
	http.ListenAndServe(":8080", nil)
}
