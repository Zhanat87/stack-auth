package routers

import (
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/controllers"
)

func SetIndexRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/test", controllers.Test).Methods("GET")
	router.HandleFunc("/test2", controllers.Test2).Methods("GET")
	return router
}