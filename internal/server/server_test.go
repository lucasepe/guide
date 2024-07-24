package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lucasepe/guide/internal/handlers"
)

func TestServer(t *testing.T) {
	exp := handlers.Greeting{
		Message: "Hello, World!",
	}

	server := httptest.NewServer(New().Handler)
	defer server.Close()

	resp, err := http.Get(server.URL + "/greet")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Leggi il corpo della risposta
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	got := handlers.Greeting{}
	if err := json.Unmarshal(dat, &got); err != nil {
		t.Errorf("Error: %v", err)
	}

	if !cmp.Equal(got, exp) {
		t.Errorf("got: %v, expected: %v", got, exp)
	}
}
