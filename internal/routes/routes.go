// internal/routes/routes.go
package routes

import (
	"net/http"

	"balzac/auth"
	"balzac/config"
	"balzac/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	// Initialize auth handler
	authHandler := auth.NewHandler(cfg)

	// CORS middleware
	r.Use(corsMiddleware(cfg.FrontendURL))

	// Auth routes
	r.HandleFunc("/auth/google/login", authHandler.GoogleLogin).Methods("GET")
	r.HandleFunc("/auth/google/callback", authHandler.GoogleCallback).Methods("GET")
	r.HandleFunc("/auth/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/auth/me", authHandler.GetUser).Methods("GET")

	// Protected API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/protected", authHandler.RequireAuth(handlers.ProtectedHandler)).Methods("GET")
	api.HandleFunc("/profile", authHandler.RequireAuth(handlers.ProfileHandler)).Methods("GET")

	// Public API routes
	r.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")

	return r
}

func corsMiddleware(frontendURL string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", frontendURL)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
