package handler

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware records the time taken for a request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
