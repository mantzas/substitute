package middleware

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mantzas/adaptlog"
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

// DefaultMiddleware which handles Logging and Recover middleware
func DefaultMiddleware(next httprouter.Handle) httprouter.Handle {
	return LoggingMiddleware(RecoveryMiddleware(next))
}

// LoggingMiddleware for recovering from failed requests
func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		lw := &statusLoggingResponseWriter{-1, false, w}
		startTime := time.Now()
		next(lw, r, ps)
		adaptlog.Printf("host=%s method=%s route=%s status=%d time=%s params=%s", r.Host, r.Method, r.URL.String(), lw.status, time.Since(startTime), ps)
	}
}

// RecoveryMiddleware for recovering from failed requests
func RecoveryMiddleware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		defer func() {
			if err := recover(); err != nil {
				adaptlog.Printf("[ERROR] %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next(w, r, ps)
	}
}
