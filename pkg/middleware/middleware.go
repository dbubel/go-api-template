// Package middleware provides HTTP middleware functions for the API server.
package middleware

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/dbubel/intake/v2"
	"github.com/sirupsen/logrus"
)

// Middleware provides HTTP middleware functions.
type Middleware struct {
	log *logrus.Logger
}

// NewMiddleware creates a new Middleware instance with the given logger.
func NewMiddleware(log *logrus.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

// Recover is a middleware that recovers from panics and logs the error.
func (m *Middleware) Recover(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				m.log.WithFields(logrus.Fields{
					"error": err,
					"stack": string(stack),
				}).Error("Panic recovered")

				// Return a 500 Internal Server Error
				intake.RespondJSON(w, r, http.StatusInternalServerError, map[string]string{
					"error": "An internal server error occurred",
				})
			}
		}()
		next(w, r)
	}
}

// Logging is a middleware that logs HTTP requests.
func (m *Middleware) Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a custom response writer to capture the status code
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Execute the next handler
		next(rw, r)

		// Log the request
		duration := time.Since(start)
		m.log.WithFields(logrus.Fields{
			"method":     r.Method,
			"path":       r.URL.Path,
			"status":     rw.statusCode,
			"duration":   duration.String(),
			"user_agent": r.UserAgent(),
			"remote_ip":  r.RemoteAddr,
		}).Info("Request processed")
	}
}

// Timeout is a middleware that applies a timeout to request processing.
func (m *Middleware) Timeout(timeout time.Duration) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			r = r.WithContext(ctx)

			next(w, r)
		}
	}
}

// CORS is a middleware that adds Cross-Origin Resource Sharing headers.
func (m *Middleware) CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// responseWriter is a custom response writer that captures the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and passes it to the underlying ResponseWriter.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures the status code (if not already set) and passes the data to the underlying ResponseWriter.
func (rw *responseWriter) Write(data []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(data)
} 