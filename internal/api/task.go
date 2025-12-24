package api

import (
	"encoding/json"
	"net/http"
	"time"

	"go.mattglei.ch/tlcore/pkg/timelog"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Title       string  `json:"title"`
		Description *string `json:"description"`
		GroupID     *uint   `json:"group_id"`

		DueTime *time.Time `json:"due_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if payload.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	if payload.GroupID == nil {
		http.Error(w, "group id is required", http.StatusBadRequest)
		return
	}

	task := timelog.Task{
		Title:       payload.Title,
		Description: payload.Description,
		GroupID:     *payload.GroupID,
		DueTime:     payload.DueTime,
		CreatedTime: time.Now(),
	}

	err = h.DB.WithContext(r.Context()).Create(&task).Error
	if err != nil {
		internalError(w, err, "failed to create task")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	h.deleteByID(w, r, "task", &timelog.Task{})
}
