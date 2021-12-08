package routes

import (
	"encoding/json"
	"net/http"
)

type QnAError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, status int, text, message string) {
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(QnAError{text, message})
}
