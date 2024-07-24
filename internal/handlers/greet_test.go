package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGreetWithName(t *testing.T) {
	exp := Greeting{
		Message: "Hello, Krateo!",
	}

	req := httptest.NewRequest(http.MethodGet, "/greet?name=Krateo", nil)

	w := httptest.NewRecorder()
	Greet(w, req)

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

func TestGreetWithNoName(t *testing.T) {
	exp := Greeting{
		Message: "Hello, World!",
	}

	req := httptest.NewRequest(http.MethodGet, "/greet", nil)

	w := httptest.NewRecorder()
	Greet(w, req)

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
