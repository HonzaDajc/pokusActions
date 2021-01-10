package main

import "testing"

func TestHelloName(t *testing.T) {
	name := "jD"
	res, _ := Hello(name)
	if res != "Hi jD" {
		t.Fatalf("Bad greet result: %s" + res)
	}
}
