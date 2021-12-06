package handlers

import (
	"encoding/json"
	"net/http"
)

func NewHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": "ok", "data": "API is running",
	})
	if err != nil {
		return
	}
}
