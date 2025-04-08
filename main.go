package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	twitterscraper "github.com/imperatrona/twitter-scraper"
	"github.com/joho/godotenv"
)

type TwitterClient struct {
	scraper *twitterscraper.Scraper
}

func NewTwitterClient() *TwitterClient {
	scraper := twitterscraper.New()
	scraper.SetAuthToken(twitterscraper.AuthToken{
		Token:     os.Getenv("TWITTER_TOKEN"),
		CSRFToken: os.Getenv("TWITTER_CSRF_TOKEN"),
	})

	if !scraper.IsLoggedIn() {
		panic("Invalid AuthToken")
	}
	return &TwitterClient{scraper: scraper}
}

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// API Handlers

// profileHandler returns the profile information of a Twitter user.
// @Summary Get Twitter profile
// @Description Get the profile information of a Twitter user by username
// @Tags Twitter
// @Param username path string true "Twitter username"
// @Success 200 {object} twitterscraper.Profile
// @Failure 400 {object} map[string]string "Invalid username"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /profile/{username} [get]
func (tc *TwitterClient) profileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, `{"error": "username is required"}`, http.StatusBadRequest)
		return
	}

	profile, err := tc.scraper.GetProfile(username)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// tweetsHandler returns the latest tweets of a Twitter user.
// @Summary Get Twitter tweets
// @Description Get the latest tweets of a Twitter user by username
// @Tags Twitter
// @Param username path string true "Twitter username"
// @Success 200 {array} twitterscraper.Tweet
// @Failure 400 {object} map[string]string "Invalid username"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /tweets/{username} [get]
func (tc *TwitterClient) tweetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		http.Error(w, `{"error": "username is required"}`, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tweets := tc.scraper.GetTweets(ctx, username, 10)
	var result []*twitterscraper.Tweet

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

func main() {
	if os.Getenv("TWITTER_TOKEN") == "" || os.Getenv("TWITTER_CSRF_TOKEN") == "" {
		log.Fatal("Missing required environment variables")
	}

	client := NewTwitterClient()

	r := mux.NewRouter()

	// API Routes
	r.HandleFunc("/profile/{username}", client.profileHandler).Methods("GET")
	r.HandleFunc("/tweets/{username}", client.tweetsHandler).Methods("GET")

	// Middleware for security and logging
	r.Use(loggingMiddleware)
	r.Use(rateLimitMiddleware)

	// Serve Swagger UI for API documentation
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./swagger-ui"))))

	http.Handle("/", r)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Middleware for rate limiting
func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement rate limiting logic here (e.g., using a token bucket or sliding window)
		next.ServeHTTP(w, r)
	})
}
