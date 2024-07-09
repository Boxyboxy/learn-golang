package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	router *http.ServeMux // returns ServerMux which is also a http.Handler
}

func NewPlayerServer(store PlayerStore) ( p *PlayerServer ){
	p = &PlayerServer{ store, http.NewServeMux()}
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	

	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return // ServerHTTP method is part of the http.Handler interface
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {	
	router := http.NewServeMux() 

	
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")
		switch r.Method {
		case http.MethodPost:
			p.processWin(w, player)
		case http.MethodGet:
			p.showScore(w, player)
		}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}	
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted) 
}