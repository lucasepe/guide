package server

import (
	"net/http"

	"github.com/lucasepe/guide/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", handlers.Greet)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
