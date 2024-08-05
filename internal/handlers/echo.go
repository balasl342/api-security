package handlers

import (
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		http.Error(w, "Message query parameter is required", http.StatusBadRequest)
		return
	}
	w.Write([]byte(message))
}
