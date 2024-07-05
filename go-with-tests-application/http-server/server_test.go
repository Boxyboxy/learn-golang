package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// implementation of the interface
type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		// mocking request
		request := newGetScoreRequest("Pepper")
		// spy made for us called ResponseRecorder to inspect what has been written as a response
		response := httptest.NewRecorder()
		// handler
		server.ServeHTTP(response, request)
		got:= response.Body.String()
		

		assertResponseBody(t, got, "20")
	})


	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		got := response.Body.String()
		
		assertResponseBody(t, got, "10")
	})
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