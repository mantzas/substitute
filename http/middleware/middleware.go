package middleware

import (
	"net/http"
	"time"

	log "github.com/mantzas/adaptlog"
)

type statusLoggingResponseWriter struct {
	status              int
	statusHeaderWritten bool
	w                   http.ResponseWriter
}

func (w *statusLoggingResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *statusLoggingResponseWriter) Write(d []byte) (int, error) {

	value, err := w.w.Write(d)
	if err != nil {
		return value, err
	}

	if !w.statusHeaderWritten {
		w.status = http.StatusOK
		w.statusHeaderWritten = true
	}

	return value, err
}

func (w *statusLoggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.w.WriteHeader(code)
	w.statusHeaderWritten = true
}

// DefaultPostMiddleware which handles Logging, POST and Recover middleware
func DefaultPostMiddleware(next http.Handler) http.Handler {
	return LoggingMiddleware(RecoveryMiddleware(PostValidationMiddleware(next)))
}

// DefaultGetMiddleware which handles Logging, GET and Recover middleware
func DefaultGetMiddleware(next http.Handler) http.Handler {
	return LoggingMiddleware(RecoveryMiddleware(GetValidationMiddleware(next)))
}

// LoggingMiddleware for logging requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &statusLoggingResponseWriter{-1, false, w}
		startTime := time.Now()
		next.ServeHTTP(lw, r)
		log.Logger.Printf("host=%s method=%s route=%s status=%d time=%s", r.Host, r.Method, r.URL.String(), lw.status, time.Since(startTime))
	})
}

// RecoveryMiddleware for recovering from failed requests
func RecoveryMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				log.Logger.Printf("[ERROR] %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// PostValidationMiddleware for validating POST requests
func PostValidationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {

			log.Logger.Printf("[WARN] Http method POST was expected, but received %s instead", r.Method)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetValidationMiddleware for validating GET requests
func GetValidationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {

			log.Logger.Printf("[WARN] Http method GET was expected, but received %s instead", r.Method)
			log.Logger.Println()
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
