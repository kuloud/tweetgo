package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kuloud/tweetgo/internal/api"
	"github.com/kuloud/tweetgo/internal/middleware"
	"github.com/kuloud/tweetgo/internal/service"

	_ "github.com/kuloud/tweetgo/docs" // Import generated docs
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("TWITTER_TOKEN") == "" || os.Getenv("TWITTER_CSRF_TOKEN") == "" {
		log.Fatal("Missing required environment variables")
	}

	twitterService := service.NewTwitterService()
	apiHandler := api.NewAPIHandler(twitterService)

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// API Routes
	apiRouter.HandleFunc("/profile/{username}", apiHandler.ProfileHandler).Methods("GET")
	apiRouter.HandleFunc("/tweets/{username}", apiHandler.TweetsHandler).Methods("GET")

	// Middleware for security and logging
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RateLimitMiddleware)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
