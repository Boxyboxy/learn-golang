package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league League
}

// refactor for constructor
func intialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err!=nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size()==0 {
		file.Write([]byte("[]"))
		file.Seek(0,io.SeekStart)
	}
	return nil
}

// constructor 
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := intialisePlayerDBFile(file)

	if err!=nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err!=nil {

		return nil,  fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{database: json.NewEncoder(&tape{file}), league:league, }, nil

}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	
	player  := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}

	// for i, player := range league {
	// 	if player.Name == name {
	// 		league[i].Wins++ // When ranging over the slice, we get copy of the element at that index.
	// 		// Changing the Wins value of the copy won't have any effect on the league slice
	// 	}
	// }

	
	f.database.Encode(f.league)

}
