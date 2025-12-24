package api

import (
	"encoding/json"
	"net/http"

	"go.mattglei.ch/tlcore/pkg/timelog"
)

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"string"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if payload.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	group := timelog.Group{Name: payload.Name}
	err = h.DB.WithContext(r.Context()).Create(&group).Error
	if err != nil {
		internalError(w, err, "failed to create group")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
