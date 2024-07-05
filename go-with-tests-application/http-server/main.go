package main

import (
	"log"
	"net/http"
)

// Complete the scaffolding

// we want to wire this up into an application because:
// 1. Actual working software, good to see the code in action
// 2. Make sure this is reflected in oru application too as part of the incremental approach

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	// Handler func is an adaptor to allow the use of ordinary functions as HTTP handlers
	// handler:= http.HandlerFunc(PlayerServer) // type casting our Player Server function with it, we have no implemented the required Handler
	// log.Fatal(http.ListenAndServe(":5000", handler))
	server := &PlayerServer{&InMemoryPlayerStore{}} //won't compile without a store
	log.Fatal(http.ListenAndServe(":5000", server))
}



// to run this, go build and then execute the exe output file