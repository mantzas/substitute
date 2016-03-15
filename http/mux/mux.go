package mux

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mantzas/substitute/http/middleware"
)

// GetServerMux creates the default request mux
func GetServerMux() *httprouter.Router {

	router := httprouter.New()
	router.GET("/*any", middleware.DefaultMiddleware(anyHandle))
	router.POST("/*any", middleware.DefaultMiddleware(anyHandle))
	router.DELETE("/*any", middleware.DefaultMiddleware(anyHandle))
	router.PUT("/*any", middleware.DefaultMiddleware(anyHandle))
	router.HEAD("/*any", middleware.DefaultMiddleware(anyHandle))
	router.OPTIONS("/*any", middleware.DefaultMiddleware(anyHandle))
	router.PATCH("/*any", middleware.DefaultMiddleware(anyHandle))
	return router
}

func anyHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("handler2"))
}
