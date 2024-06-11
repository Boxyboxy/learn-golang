package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want) // %q prints the value as double-quoted string
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want) // %q prints the value as double-quoted string
	}
}

// go mod init hello
// create go.mod file
// This file tells go tools essential information about my code
// writing tests:
/*
1.file name: xxx_test.go
2. test function must start with the word Test
3. test function takes one argument only t *testing.T
4. Need to import "testing"
*/

// go test -v
