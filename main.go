package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type EnvVarsModel struct {
	vars []string `json:"vars"`
	mongoDbHost string `json:"mongoDbHost"`
	a string `json:"a"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		output, err := json.Marshal(EnvVarsModel{vars: os.Environ(), mongoDbHost: os.Getenv("STACK_MONGODB_PORT_27017_TCP_ADDR"), a: "b"})
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)

	})
	http.ListenAndServe(":8082", nil)
}