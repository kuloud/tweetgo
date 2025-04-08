package api

import (
	"net/http"
	"os"

	"github.com/kuloud/tweetgo/pkg/auth"
)

func (h *APIHandler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<h1>Welcome to the X Scraper API</h1>
<p>Visit <a href="/swagger/">Swagger UI</a> to explore the API.</p>`))
}

// LoginHandler authenticates a user and returns a JWT token.
// @Summary Login to get a JWT token
// @Description Authenticate a user and return a JWT token
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} map[string]string "Returns a JWT token"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /login [post]
func (h *APIHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// TODO: Implement authentication logic here
	if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(username)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token": "` + token + `"}`))
}
