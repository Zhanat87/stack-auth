package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Zhanat87/stack-auth/common"
)

type EnvVarsModel struct {
	vars []string `json:"vars"`
	mongoDbHost string `json:"mongoDbHost"`
}

// Handler for HTTP Get - "/"
func Index(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(EnvVarsModel{vars: os.Environ(), mongoDbHost: os.Getenv("STACK_MONGODB_PORT_27017_TCP_ADDR")})
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