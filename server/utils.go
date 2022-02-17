package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type RequestDataPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func getKeyFromURL(url string) (string, error) {
	paths := strings.Split(url, "/")
	if len(paths) < 3 {
		return "", errors.New("key is not found in path")
	}
	return paths[2], nil
}

func parsePayload(w http.ResponseWriter, r *http.Request) (*RequestDataPayload, error) {
	payload := &RequestDataPayload{}

	err := r.ParseForm()
	if err != nil {
		return nil, NewClientError(http.StatusBadRequest, err)
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, NewClientError(http.StatusBadRequest, err)
	}

	return payload, nil
}
