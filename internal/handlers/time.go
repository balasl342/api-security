package handlers

import (
	"net/http"
	"time"
)

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format(time.RFC3339)))
}
