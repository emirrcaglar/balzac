package handlers

import (
	"balzac/auth"
	"encoding/json"
	"net/http"
)

// Example handlers
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserFromContext(r.Context())

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "This is a protected endpoint",
		"user":    user,
	}
	json.NewEncoder(w).Encode(response)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserFromContext(r.Context())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
