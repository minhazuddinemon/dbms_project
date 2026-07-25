package logger

import (
	"log"
	"net/http"
	"time"
)

// RequestLogger is a middleware that logs all incoming HTTP requests
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass the request to the next handler
		next.ServeHTTP(w, r)

		// Log the result after the handler finishes
		log.Printf(
			"[%s] %s | %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
