package main

import (
	"github.com/lucasepe/guide/internal/server"
)

func main() {
	server := server.New()
	server.ListenAndServe()
}
