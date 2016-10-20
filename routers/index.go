package routers
import (
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/controllers"
)
func SetIndexRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/index", controllers.Index).Methods("GET")
	return router
}