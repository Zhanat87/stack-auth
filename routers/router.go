package routers
import (
	"github.com/gorilla/mux"
	"github.com/Zhanat87/stack-auth/controllers"
)
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/test", controllers.GetUsers).Methods("GET")
	// Routes for the index
	router = SetIndexRoutes(router)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Task entity
	router = SetTaskRoutes(router)
	// Routes for the TaskNote entity
	router = SetNoteRoutes(router)
	return router
}