package routers

import (
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/controllers"
)

func SetIndexRoutes(router *mux.Router) *mux.Router {
	indexRouter := mux.NewRouter()
	indexRouter.HandleFunc("/index", controllers.Index).Methods("GET")
	indexRouter.HandleFunc("/test", controllers.Test).Methods("GET")
	return indexRouter
}