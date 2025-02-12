package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		message = "Hello, World!"
	}

	testEnv := os.Getenv("TEST")
	if testEnv == "" {
		testEnv = "Environment variable TEST is not set"
	}

	log.Printf("Received request with message: %s, TEST: %s", message, testEnv)
	fmt.Fprintf(w, "Echo: %s\nTEST: %s\n", message, testEnv)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/health", healthHandler)
	log.Println("Starting echo server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}