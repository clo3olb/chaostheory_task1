package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/clo3olb/chaostheory_task1/database"
)

func listHandler(w http.ResponseWriter, r *http.Request) error {
	err := json.NewEncoder(w).Encode(database.DB().GetDataList())
	if err != nil {
		return err
	}
	return nil
}

func addHandler(w http.ResponseWriter, r *http.Request) error {

	payload, err := parsePayload(w, r)
	if err != nil {
		return err
	}

	err = database.DB().Create(payload.Key, payload.Value)
	if err != nil {
		errorMessage := "data already exits"
		return NewClientError(http.StatusConflict, errors.New(errorMessage))
	}

	responseMessage := "Data created."
	json.NewEncoder(w).Encode(response{responseMessage})
	w.WriteHeader(http.StatusCreated)
	return nil
}

func createDataHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getKeyFromURL(r.URL.Path)
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	payload, err := parsePayload(w, r)
	if err != nil {
		return err
	}

	err = database.DB().Create(key, payload.Value)
	if err != nil {
		return NewClientError(http.StatusConflict, err)
	}

	responseMessage := "Data created."
	json.NewEncoder(w).Encode(response{responseMessage})
	w.WriteHeader(http.StatusCreated)
	return nil
}

func readDataHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getKeyFromURL(r.URL.Path)
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	data, err := database.DB().Find(key)
	if err != nil {
		return NewClientError(http.StatusNotFound, err)
	}

	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusOK)
	return nil
}

func updateDataHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getKeyFromURL(r.URL.Path)
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	payload, err := parsePayload(w, r)
	if err != nil {
		return err
	}

	err = database.DB().Update(key, payload.Value)
	if err != nil {
		return NewClientError(http.StatusNotFound, err)
	}

	responseMessage := "Data updated."
	json.NewEncoder(w).Encode(response{responseMessage})
	w.WriteHeader(http.StatusOK)
	return nil
}

func deleteDataHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getKeyFromURL(r.URL.Path)
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	err = database.DB().Delete(key)
	if err != nil {
		return NewClientError(http.StatusNotFound, err)
	}
	json.NewEncoder(w).Encode(response{"Data deleted."})
	w.WriteHeader(http.StatusOK)

	return nil
}
