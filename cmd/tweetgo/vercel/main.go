package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kuloud/tweetgo/pkg/api"
	"github.com/kuloud/tweetgo/pkg/middleware"
	"github.com/kuloud/tweetgo/pkg/service"
	"github.com/rs/cors"

	_ "github.com/kuloud/tweetgo/docs" // Import generated docs
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// @title X Scraper API
// @version 1.0
// @description API for scraping X profiles and tweets
// @host tweetgo.vercel.app
// @BasePath /
func Handler_cmd_tweetgo_main_go_cmd_tweetgo_vercel_main_go(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("TWITTER_TOKEN") == "" || os.Getenv("TWITTER_CSRF_TOKEN") == "" {
		log.Fatal("Missing required environment variables")
	}

	twitterService := service.NewTwitterService()
	apiHandler := api.NewAPIHandler(twitterService)

	router := mux.NewRouter()

	router.HandleFunc("/login", apiHandler.LoginHandler).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware)

	// API Routes
	apiRouter.HandleFunc("/profile/{username}", apiHandler.ProfileHandler).Methods("GET")
	apiRouter.HandleFunc("/tweets/{username}", apiHandler.TweetsHandler).Methods("GET")

	// Middleware for security and logging
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RateLimitMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	c.Handler(router)

	router.ServeHTTP(w, r)
}
