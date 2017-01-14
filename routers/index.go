package routers

import (
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/controllers"
)

func SetIndexRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/index", controllers.Index).Methods("GET")
	router.HandleFunc("/test", controllers.Test).Methods("GET")
	return router
}