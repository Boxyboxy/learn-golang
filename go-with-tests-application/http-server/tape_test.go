package main

import (
	"io"
	"os"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, cleanDatabase := createTempFile(t, "12345")
	defer cleanDatabase()

	tape := &tape{file.(*os.File)}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)

	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}