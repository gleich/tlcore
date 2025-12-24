package api

import (
	"fmt"
	"net/http"
	"strconv"

	"go.mattglei.ch/timber"
)

func (h *Handler) deleteByID(w http.ResponseWriter, r *http.Request, name string, model any) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, fmt.Sprintf("missing %s id", name), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid %s id", name), http.StatusBadRequest)
		return
	}

	tx := h.DB.WithContext(r.Context()).Delete(model, id)
	err = tx.Error
	if err != nil {
		internalError(w, err, fmt.Sprintf("failed to delete %s", name))
		return
	}

	if tx.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("%s not found", name), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func internalError(w http.ResponseWriter, err error, msg string) {
	http.Error(w, "internal server occurred", http.StatusInternalServerError)
	timber.Error(err, msg)
}
