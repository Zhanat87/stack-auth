package routers
import (
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/controllers"
)
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/register", controllers.Register).Methods("GET")
	router.HandleFunc("/users/login", controllers.Login).Methods("GET")
	return router
}