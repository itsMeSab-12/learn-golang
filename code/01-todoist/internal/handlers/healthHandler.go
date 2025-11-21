package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode("health check ok")
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
