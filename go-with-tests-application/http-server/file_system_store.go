package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	}
	// for i, player := range league {
	// 	if player.Name == name {
	// 		league[i].Wins++ // When ranging over the slice, we get copy of the element at that index.
	// 		// Changing the Wins value of the copy won't have any effect on the league slice
	// 	}
	// }

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}
