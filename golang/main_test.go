package main

import "testing"

func TestHelloWorld(t *testing.T) {

	expected := "Hello from Golang!"
	result := HelloWorld()

	if expected != result {
		t.Errorf("HelloWorld was incorrect, got: %s, want: %s.", result, expected)
	}
}
