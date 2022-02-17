package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	port          int    = 80
	serverBaseUrl string = "http://localhost"
)

func Start() {

	// Creates Router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// page handlers
	r.Get("/", handle(documentationHandler))
	r.Get("/list", handle(listHandler))
	r.Post("/add", handle(addHandler))

	// Logging server starts.
	fmt.Printf("Server is listening: %s:%d\n", serverBaseUrl, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
