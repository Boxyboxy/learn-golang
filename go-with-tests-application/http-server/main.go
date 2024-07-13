package main

import (
	"log"
	"net/http"
	"os"
)

// Complete the scaffolding

// we want to wire this up into an application because:
// 1. Actual working software, good to see the code in action
// 2. Make sure this is reflected in oru application too as part of the incremental approach



const dbFileName = "game.db.json"
func main() {
	// Handler func is an adaptor to allow the use of ordinary functions as HTTP handlers
	// handler:= http.HandlerFunc(PlayerServer) // type casting our Player Server function with it, we have no implemented the required Handler
	// log.Fatal(http.ListenAndServe(":5000", handler))
	//store := NewInMemoryPlayerStore()


	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s: %v", dbFileName, err)
	}
	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("didnt expect an error but got one, %v",  err)
	}
	server := NewPlayerServer(store) //won't compile without a store
	log.Fatal(http.ListenAndServe(":5000", server))
}

// to run this, go build and then execute the exe output file