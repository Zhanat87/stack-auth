package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Zhanat87/stack-auth/common"
)

type EnvVarsModel struct {
	Vars []string `json:"vars"`
	MongoDbHost string `json:"mongoDbHost"`
}

type EnvVariablesModel struct {
	Vars []string `json:"vars"`
	A string `json:"a"`
	C int `json:"c"`
}

type EnvTestModel struct {
	A string `json:"a"`
	C int `json:"c"`
}

// Handler for HTTP Get - "/"
func Index(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(EnvVarsModel{Vars: os.Environ(), MongoDbHost: os.Getenv("STACK_MONGODB_PORT_27017_TCP_ADDR")})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Handler for HTTP Get - "/test"
func Test(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(EnvVariablesModel{Vars: os.Environ(), A: "b", C: 123})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Handler for HTTP Get - "/test2"
func Test2(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(EnvTestModel{A: "b", C: 123})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}