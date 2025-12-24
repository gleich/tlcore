package api

import (
	"net/http"

	"go.mattglei.ch/timber"
)

func internalError(w http.ResponseWriter, err error, msg string) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	timber.Error(err, msg)
}
