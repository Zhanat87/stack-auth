package routers
import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/Zhanat87/stack-auth/common"
	"github.com/Zhanat87/stack-auth/controllers"
)
func SetUserRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	userRouter.HandleFunc("/sign-up", controllers.CreateUser).Methods("POST")
	
	userRouter.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	userRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.PathPrefix("/users").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(userRouter),
	))
	return router
}