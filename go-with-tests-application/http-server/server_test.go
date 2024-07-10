package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// implementation of the interface
type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
	league []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
    return s.league
}

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		// mocking request
		request := newGetScoreRequest("Pepper")
		// spy made for us called ResponseRecorder to inspect what has been written as a response
		response := httptest.NewRecorder()
		// handler
		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})


	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		got := response.Code
		want := http.StatusNotFound
		
		assertStatus(t, got, want)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}


func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name ), nil)
	return request
}


func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func TestStoreWins(t *testing.T){
	store:= StubPlayerStore{
		map[string]int{},	
		nil,nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestLeague(t *testing.T) {

	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}

	store := StubPlayerStore{nil, nil, wantedLeague}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		
		var got []Player = getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)

		assertContentType(t, response, jsonContentType)
		
	})

	
}

func assertLeague(t testing.TB, got, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague){
		t.Errorf("got %v, want %v", got, wantedLeague)        
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string){
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %v, got %v", want, response.Result().Header.Get("content-type"))
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player){
	t.Helper()
	// create a decoder from encoding/json package
	// it needs an io.Reader to read from which in our case is our response spy's body
	// to create an encoder, you need an io.reader which is what httptest.ResponseRecorder.Body implements
		
	err := json.NewDecoder(body).Decode(&league)

	if err!= nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}