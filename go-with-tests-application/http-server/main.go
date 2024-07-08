package main

import (
	"log"
	"net/http"
)

// Complete the scaffolding

// we want to wire this up into an application because:
// 1. Actual working software, good to see the code in action
// 2. Make sure this is reflected in oru application too as part of the incremental approach




func main() {
	// Handler func is an adaptor to allow the use of ordinary functions as HTTP handlers
	// handler:= http.HandlerFunc(PlayerServer) // type casting our Player Server function with it, we have no implemented the required Handler
	// log.Fatal(http.ListenAndServe(":5000", handler))
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store} //won't compile without a store
	log.Fatal(http.ListenAndServe(":5000", server))
}

// to run this, go build and then execute the exe output file