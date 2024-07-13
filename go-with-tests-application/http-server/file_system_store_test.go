package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"name": "Cleo", "wins": 10},
			{"name": "Chris", "wins": 33}
		]`)
		defer cleanDatabase() // ensure that a function call is performed later in a program's executioni, usually for purpose of clean up

		store, err := NewFileSystemPlayerStore( database.(*os.File))
		assertNoError(t, err)
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	
		// read again. Will only work if the store implements ReadSeeker/ ReadWriteSeeker
		// Prior: When using Reader, Reader has reached the end so there is nothing more to read. We need a way to tell it to go back to the start. 
		got = store.GetLeague()
		assertLeague(t, got, want)
	})


	t.Run("get player score", func(t *testing.T) {
		
		database, cleanDatabase := createTempFile(t, `[
			{"name": "Cleo", "wins": 10},
			{"name": "Chris", "wins": 33}
		]`)
		defer cleanDatabase() // ensure that a function call is performed later in a program's executioni, usually for purpose of clean up

		store, err:= NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	//file_system_store_test.go
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database.(*os.File))

		assertNoError(t, err)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T){
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database.(*os.File))
		assertNoError(t, err)
	})
}

func createTempFile( t testing.TB, initialData string )(io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db") // create temp file with prefix "db". prefix will be appended to random file name

	if err!=nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))



	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	} 

	return tmpfile, removeFile

	
}
func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}


func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err!=nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}