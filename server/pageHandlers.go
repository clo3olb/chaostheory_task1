package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clo3olb/chaostheory_task1/database"
)

type pageInfo struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

func getPageUrl(path string) string {
	return fmt.Sprintf("%s:%d%s", serverBaseUrl, port, path)
}

var pages = []pageInfo{
	{
		Path:        getPageUrl("/"),
		Method:      "GET",
		Description: "Displays paths and descriptions of each endpoints",
	},
	{
		Path:        getPageUrl("/list"),
		Method:      "GET",
		Description: "Lists all the data in the database as an array",
	},
	{
		Path:        getPageUrl("/add"),
		Method:      "POST",
		Description: "Adds data to the database. Example Format : { \"key\": \"string\", \"value\": \"string\" }",
	},
}

func documentationHandler(w http.ResponseWriter, r *http.Request) error {
	err := json.NewEncoder(w).Encode(pages)
	if err != nil {
		return err
	}
	return nil
}

func listHandler(w http.ResponseWriter, r *http.Request) error {
	err := json.NewEncoder(w).Encode(database.GetDataList(database.DB()))
	if err != nil {
		return err
	}
	return nil
}

func addHandler(w http.ResponseWriter, r *http.Request) error {
	payload := database.Data{}

	err := r.ParseForm()
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	newData, err := database.FormatData(payload.Key, payload.Value)
	if err != nil {
		return NewClientError(http.StatusBadRequest, err)
	}

	responseMessage := "Data created."
	if database.DB().Exists(newData.Key) {
		responseMessage = "Data changed."
	}

	database.DB().Add(*newData)
	json.NewEncoder(w).Encode(response{responseMessage})
	w.WriteHeader(http.StatusCreated)
	return nil
}
