// internal/auth/handlers.go
package auth

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"balzac/config"

	"github.com/gorilla/sessions"
)

type Handler struct {
	service *Service
	store   *sessions.CookieStore
	config  *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	service := NewService(cfg)
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))

	gob.Register(&User{})

	// Configure session options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}

	return &Handler{
		service: service,
		store:   store,
		config:  cfg,
	}
}

func (h *Handler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url, state, err := h.service.GenerateAuthURL()
	if err != nil {
		log.Printf("Error generating auth URL: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set state cookie
	cookie := &http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   false, // Set to true in production
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Verify state
	oauthState, err := r.Cookie("oauthstate")
	if err != nil {
		log.Printf("No oauth state cookie: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if r.FormValue("state") != oauthState.Value {
		log.Printf("Invalid oauth state")
		http.Error(w, "Invalid oauth state", http.StatusBadRequest)
		return
	}

	// Exchange code for token
	token, err := h.service.ExchangeCodeForToken(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Get user info
	user, err := h.service.GetUserInfo(r.Context(), token)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	// Save user to session
	session, err := h.store.Get(r, "session")
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	session.Values["user"] = user
	if err := session.Save(r, w); err != nil {
		log.Printf("Failed to save session: %v", err)
		http.Error(w, fmt.Sprintf("Failed to save session: %v", err), http.StatusInternalServerError)
		return
	}

	// Clear state cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "oauthstate",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	// Redirect to frontend
	http.Redirect(w, r, h.config.FrontendURL+"/dashboard", http.StatusTemporaryRedirect)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, "session")
	if err != nil {
		log.Printf("Failed to get session: %v", err)
	}

	session.Values["user"] = nil
	session.Options.MaxAge = -1
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{
		Message: "Logged out successfully",
	})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, "session")
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	user, exists := session.Values["user"]
	if !exists || user == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(AuthResponse{
			Error: "Not authenticated",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{
		User: user.(*User),
	})
}

func (h *Handler) GetCurrentUser(r *http.Request) *User {
	session, err := h.store.Get(r, "session")
	if err != nil {
		return nil
	}

	user, exists := session.Values["user"]
	if !exists || user == nil {
		return nil
	}

	return user.(*User)
}
