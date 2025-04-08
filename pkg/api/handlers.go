package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	xScraper "github.com/imperatrona/twitter-scraper"
	"github.com/kuloud/tweetgo/pkg/service"
)

// @title X Scraper API
// @version 1.0
// @description API for scraping X profiles and tweets
// @BasePath /api/v1
type APIHandler struct {
	xService *service.XService
}

func NewAPIHandler(service *service.XService) *APIHandler {
	return &APIHandler{xService: service}
}

// GetXIdByUsernameHandler returns the ID of an X user by username.
// @Summary Get X ID by username
// @Description Get the ID of an X user by username
// @Tags XRating
// @Param username query string true "X username"
// @Success 200 {object} map[string]string "Returns the X ID"
// @Failure 400 {object} map[string]string "Invalid username"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/xrating/getXIdByUsername [get]
func (h *APIHandler) GetXIdByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, `{"error": "username is required"}`, http.StatusBadRequest)
		return
	}

	// Call the service to get the X ID
	id, err := h.xService.GetXIdByUsername(username)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// ProfileHandler returns the profile information of a X user.
// @Summary Get X profile
// @Description Get the profile information of a X user by username
// @Tags X
// @Param username path string true "X username"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} map[string]string "Profile information xScraper.Profile"
// @Failure 400 {object} map[string]string "Invalid username"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/profile/{username} [get]
func (h *APIHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, `{"error": "username is required"}`, http.StatusBadRequest)
		return
	}

	profile, err := h.xService.GetProfile(username)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// TweetsHandler returns the latest tweets of a X user.
// @Summary Get X tweets
// @Description Get the latest tweets of a X user by username
// @Tags X
// @Param username path string true "X username"
// @Param Authorization header string true "Bearer token"
// @Success 200 {array}  []map[string]string "Latest tweets []xScraper.Tweet"
// @Failure 400 {object} map[string]string "Invalid username"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/tweets/{username} [get]
func (h *APIHandler) TweetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, `{"error": "username is required"}`, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tweets := h.xService.GetTweets(ctx, username, 10)
	var result []*xScraper.Tweet

	for tweet := range tweets {
		if tweet.Error != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%v"}`, tweet.Error), http.StatusInternalServerError)
			return
		}
		result = append(result, &tweet.Tweet)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
