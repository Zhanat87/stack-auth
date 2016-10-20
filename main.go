package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/Zhanat87/stack-auth/controllers"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/index", controllers.Index).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}