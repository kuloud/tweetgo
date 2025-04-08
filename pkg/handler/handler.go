package handler

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kuloud/tweetgo/pkg/api"
	"github.com/kuloud/tweetgo/pkg/middleware"
	"github.com/kuloud/tweetgo/pkg/service"
	"github.com/rs/cors"

	xScraper "github.com/imperatrona/twitter-scraper"
	_ "github.com/kuloud/tweetgo/docs" // Import generated docs
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewHandler() http.Handler {
	// Define multiple Twitter tokens in the format: TWITTER_TOKENS="token1:csrf1,token2:csrf2"
	tokensEnv := os.Getenv("TWITTER_TOKENS")
	if tokensEnv == "" {
		log.Fatal("Missing required environment variable: TWITTER_TOKENS")
	}

	// Parse tokens
	var tokens []xScraper.AuthToken
	for _, tokenPair := range strings.Split(tokensEnv, ",") {
		parts := strings.Split(tokenPair, ":")
		if len(parts) != 2 {
			log.Fatal("Invalid token format in TWITTER_TOKENS")
		}
		tokens = append(tokens, xScraper.AuthToken{
			Token:     strings.TrimSpace(parts[0]),
			CSRFToken: strings.TrimSpace(parts[1]),
		})
	}
	pool := service.NewXServicePool(tokens)
	apiHandler := api.NewAPIHandler(pool)

	router := mux.NewRouter()

	router.HandleFunc("/", apiHandler.DefaultHandler)
	router.HandleFunc("/login", apiHandler.LoginHandler).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware)

	// API Routes
	apiRouter.HandleFunc("/profile/{username}", apiHandler.ProfileHandler).Methods("GET")
	apiRouter.HandleFunc("/tweets", apiHandler.TweetsHandler).Methods("GET")
	apiRouter.HandleFunc("/tweets/{tweetId}", apiHandler.TweetHandler).Methods("GET")

	// Middleware for security and logging
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RateLimitMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	return c.Handler(router)
}
