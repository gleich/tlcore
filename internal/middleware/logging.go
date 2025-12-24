package middleware

import (
	"net/http"
	"strings"
	"time"

	"go.mattglei.ch/timber"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{ResponseWriter: w}
		next.ServeHTTP(wrapped, r)
		timber.Donef(
			"%d [%s] %s %s %s",
			wrapped.statusCode,
			strings.ToLower(http.StatusText(wrapped.statusCode)),
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
