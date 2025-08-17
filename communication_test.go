package main

import (
	"testing"
	"time"
)

func TestCommunication(t *testing.T) {
	port := "8081" // Use a different port for testing
	go runServer(port)

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	err := runClient(port)
	if err != nil {
		t.Errorf("Client returned an error: %v", err)
	}
}
