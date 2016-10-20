package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/Zhanat87/stack-auth/common"
	"github.com/Zhanat87/stack-auth/routers"
	"os"
)

//Entry point of the program
func main() {
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		//Addr:    common.AppConfig.Server,
		Addr:    os.Getenv("STACK_AUTH_PORT_8082_TCP_ADDR"),
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
