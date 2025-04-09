package handler

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kuloud/tweetgo/pkg/database"
	"github.com/kuloud/tweetgo/pkg/handler"

	_ "github.com/kuloud/tweetgo/docs" // Import generated docs
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

// @title X Scraper API
// @version 1.0
// @description API for scraping X profiles and tweets
// @host tweetgo.vercel.app
// @BasePath /
func Handler_cmd_tweetgo_main_go_cmd_tweetgo_vercel_main_go(w http.ResponseWriter, r *http.Request) {
	h := handler.NewHandler()
	h.ServeHTTP(w, r)
}
