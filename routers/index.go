package routers
import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/common"
	"github.com/Zhanat87/stack-auth/controllers"
)
func SetIndexRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/", controllers.Index).Methods("GET")
	return router
}