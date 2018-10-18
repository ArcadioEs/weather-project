package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondWithCodeAndMessage(code int, msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(code)
	response := errorResponse{code, msg}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error("Error sending response", err)
	}
}
