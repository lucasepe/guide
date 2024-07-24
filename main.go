package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lucasepe/guide/docs" // This is to import generated docs

	httpSwagger "github.com/swaggo/http-swagger"
)

// Greeting represents a greeting message
type Greeting struct {
	Message string `json:"message"`
}

// @title Hello World API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /

// greetHandler godoc
// @Summary Greet the user
// @Description get a greeting message
// @ID greet
// @Produce  json
// @Param name query string false "Name to greet"
// @Success 200 {object} main.Greeting
// @Router /greet [get]
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	response := Greeting{Message: fmt.Sprintf("Hello, %s!", name)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// NewServer creates and returns a new HTTP server with the given handler
func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greetHandler)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func main() {
	server := NewServer()
	server.ListenAndServe()
}
