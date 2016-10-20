package common

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Text string `json:"text"`
}

// need to make function exportable with an uppercase for its name:
func RespondJson(text string, w http.ResponseWriter) {
	response := Response{text}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}