package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Thomas", "French")
		want := "Bonjour, Thomas"
		assertCorrectMessage(t, got, want)
	})

}

// testing.TB is an interface that *testing.T  (test) and *testing.B (benchmark) both satisfy.
// what is the *? is it a pointer?
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tells the test suite that this method is a helper method,
	//when the test fails, the line number reported will be in our function call rather than inside our test helper.
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
