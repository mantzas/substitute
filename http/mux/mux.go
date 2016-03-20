package mux

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mantzas/substitute/http/handles"
	"github.com/mantzas/substitute/http/middleware"
)

// GetServerMux creates the default request mux
func GetServerMux() *httprouter.Router {

	router := httprouter.New()
	router.GET("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	router.POST("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	router.DELETE("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	router.PUT("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	router.HEAD("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	router.OPTIONS("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	router.PATCH("/*any", middleware.DefaultMiddleware(handles.AnyHandle))
	return router
}
