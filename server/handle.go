package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request) error

type response struct {
	Data string `json:"data"`
}

func handle(fn handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			switch e := err.(type) {
			case ClientError: // Client Error
				fmt.Printf("Client Error: %s\n", err)
				w.WriteHeader(e.StatusCode())
				json.NewEncoder(w).Encode(response{err.Error()})
			default: // Server Error
				fmt.Printf("Server Error: %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response{"Internal Server Error"})
			}
		}

	}
}
