package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lucasepe/guide/docs" // This is to import generated docs
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

// Greet godoc
// @Summary Greet the user
// @Description get a greeting message
// @ID greet
// @Produce  json
// @Param name query string false "Name to greet"
// @Success 200 {object} main.Greeting
// @Router /greet [get]
func Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	response := Greeting{Message: fmt.Sprintf("Hello, %s!", name)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
