package myjson

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, errorMessage, userMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	type error struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}

	var msg error
	msg.Error = errorMessage
	msg.Message = userMessage

	json.NewEncoder(w).Encode(msg)
	return
}
