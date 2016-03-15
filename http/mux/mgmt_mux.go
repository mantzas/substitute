package mux

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mantzas/substitute/http/middleware"
)

// GetMgmtServerMux creates the management request mux
func GetMgmtServerMux() *httprouter.Router {

	router := httprouter.New()
	router.GET("/configs", middleware.DefaultMiddleware(getConfigurationsHandler))
	return router
}

func getConfigurationsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get configs handle"))
}
