package main

import "testing"

func TestGetUUID(t *testing.T) {
	uuid := getUUID()
	if uuid == "" {
		t.Error("Error creating UUID")
	}
}
