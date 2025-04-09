package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	xScraper "github.com/imperatrona/twitter-scraper"
	_ "github.com/lib/pq"
)

var db *sql.DB
var dbEnabled bool = false

func Init() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbEnabled = false
		return nil
	}

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	dbEnabled = true
	return db.Ping()
}

func GetTweetsByUsername(ctx context.Context, username string) ([]*xScraper.Tweet, error) {
	if !dbEnabled {
		return nil, nil
	}
	rows, err := db.QueryContext(ctx, "SELECT * FROM tweets WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []*xScraper.Tweet
	for rows.Next() {
		var tweet xScraper.Tweet
		if err := rows.Scan(&tweet.ID, &tweet.Text, &tweet.Username); err != nil {
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}
	return tweets, nil
}

func GetTweetsByUserID(ctx context.Context, userID string) ([]*xScraper.Tweet, error) {
	if !dbEnabled {
		return nil, nil
	}
	rows, err := db.QueryContext(ctx, "SELECT * FROM tweets WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []*xScraper.Tweet
	for rows.Next() {
		var tweet xScraper.Tweet
		if err := rows.Scan(&tweet.ID, &tweet.Text, &tweet.Username); err != nil {
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}
	return tweets, nil
}
