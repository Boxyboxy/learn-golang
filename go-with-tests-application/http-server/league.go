package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league from JSON: %v", err)
	}
	return league, err
}

type League []Player
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i] // returns a pointer/address to the player struct
			// allows the caller to modify the Player struct directly if needed
		}
	}
	return nil
}