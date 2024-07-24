package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGreetHandlerWithName(t *testing.T) {
	exp := Greeting{
		Message: "Hello, Krateo!",
	}

	req := httptest.NewRequest(http.MethodGet, "/greet?name=Krateo", nil)

	w := httptest.NewRecorder()
	greetHandler(w, req)

	resp := w.Result()
	if resp != nil {
		defer resp.Body.Close()
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	got := Greeting{}
	if err := json.Unmarshal(dat, &got); err != nil {
		t.Errorf("Error: %v", err)
	}

	if !cmp.Equal(got, exp) {
		t.Errorf("got: %v, expected: %v", got, exp)
	}
}

func TestGreetHandlerWithNoName(t *testing.T) {
	exp := Greeting{
		Message: "Hello, World!",
	}

	req := httptest.NewRequest(http.MethodGet, "/greet", nil)

	w := httptest.NewRecorder()
	greetHandler(w, req)

	resp := w.Result()
	if resp != nil {
		defer resp.Body.Close()
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	got := Greeting{}
	if err := json.Unmarshal(dat, &got); err != nil {
		t.Errorf("Error: %v", err)
	}

	if !cmp.Equal(got, exp) {
		t.Errorf("got: %v, expected: %v", got, exp)
	}
}

func TestServer(t *testing.T) {
	exp := Greeting{
		Message: "Hello, World!",
	}

	server := httptest.NewServer(NewServer().Handler)
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

	got := Greeting{}
	if err := json.Unmarshal(dat, &got); err != nil {
		t.Errorf("Error: %v", err)
	}

	if !cmp.Equal(got, exp) {
		t.Errorf("got: %v, expected: %v", got, exp)
	}
}
