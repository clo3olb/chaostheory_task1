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

func Start() error {

	// Creates Router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// data handlers
	r.Get("/list", handle(listHandler))
	r.Post("/add", handle(addHandler))

	// CRUD approach for REST API of data
	r.Get("/data", handle(listHandler))
	r.Get("/data/{key}", handle(readDataHandler))
	r.Post("/data/{key}", handle(createDataHandler))
	r.Put("/data/{key}", handle(updateDataHandler))
	r.Delete("/data/{key}", handle(deleteDataHandler))

	// Logging server starts.
	fmt.Printf("Server is listening: %s:%d\n", serverBaseUrl, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
