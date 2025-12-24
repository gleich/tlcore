package main

import (
	"net/http"
	"time"

	"go.mattglei.ch/timber"
	"go.mattglei.ch/tlcore/internal/api/tasks"
	"go.mattglei.ch/tlcore/internal/middleware"
)

func main() {
	// setting up logging
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		timber.Fatal(err, "failed to load new york timezone")
	}
	timber.Timezone(ny)
	timber.TimeFormat("01/02 03:04:05 PM MST")

	timber.Info("booted")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/gleich/tlcore", http.StatusPermanentRedirect)
	})
	mux.HandleFunc("POST /task", tasks.Create)

	loggingMux := middleware.Logging(mux)

	timber.Info("starting server")
	server := &http.Server{
		Addr:         ":8000",
		Handler:      loggingMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		timber.Fatal(err, "failed to start router")
	}
}
