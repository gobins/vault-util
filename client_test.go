package main

import "testing"

func TestGetClient(t *testing.T) {
	client := getClient("http://localhost:8200")
	if client == nil {
		t.Error("Error creating vault client")
	}
}
