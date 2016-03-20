package handles

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mantzas/substitute/routes"
)

// AnyHandle handles all incomming request by matching it against registered and returning the stored result
func AnyHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	contentType := r.Header.Get("Content-Type")

	requestType, err := routes.ContentTypeToRequestType(contentType)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Content type not supported!"))
		return
	}

	matched, content, responseCode := routes.Register.Match(r.Method, requestType, r.URL.Path)

	if !matched {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Route not matched!"))
		return
	}

	w.WriteHeader(responseCode)
	w.Write([]byte(content))
}
