package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go.mattglei.ch/tlcore/pkg/timelog"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Title       string  `json:"title"`
		Description *string `json:"description"`
		GroupID     uint    `json:"group_id"`

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

	task := timelog.Task{
		Title:       payload.Title,
		Description: payload.Description,
		GroupID:     payload.GroupID,
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
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing task id", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	err = h.DB.WithContext(r.Context()).Delete(&timelog.Task{}, id).Error
	if err != nil {
		internalError(w, err, "failed to delete task")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
