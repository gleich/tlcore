package main

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mattglei.ch/timber"
	"go.mattglei.ch/tlcore/internal/api"
	"go.mattglei.ch/tlcore/internal/db"
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

	// load .env file for development
	err = godotenv.Load()
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		timber.Fatal(err, "loading .env file failed")
	} else {
		timber.Done("loaded .env file")
	}

	// connect to database
	database, err := db.Connect()
	if err != nil {
		timber.Fatal(err, "failed to connect to postgres database")
	}
	timber.Done("connected to postgres database")

	// run migrations (only in development)
	if os.Getenv("RUN_MIGRATIONS") == "true" {
		err = db.RunMigrations(database)
		if err != nil {
			timber.Fatal(err, "failed to run dev migrations")
		}
	}

	// register endpoints with mux
	handler := api.Handler{DB: database}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/gleich/tlcore", http.StatusPermanentRedirect)
	})
	mux.HandleFunc("POST /task", http.HandlerFunc(handler.CreateTask))
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
