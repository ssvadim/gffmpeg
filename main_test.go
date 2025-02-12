package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
)

func TestEchoHandler(t *testing.T) {
	os.Setenv("TEST", "my-test-value")

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/?message=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(echoHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v; got %v", http.StatusOK, status)
	}

	// Check the response body
	expected := "Echo: test\nTEST: my-test-value\n"
	if rr.Body.String() != expected {
		t.Errorf("Expected response: %v; got %v", expected, rr.Body.String())
	}
}
