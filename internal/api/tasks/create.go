package tasks

import (
	"encoding/json"
	"net/http"

	"go.mattglei.ch/timber"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Title       string  `json:"title"`
		Description *string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timber.Debug(payload.Title)
	w.WriteHeader(http.StatusCreated)
}
