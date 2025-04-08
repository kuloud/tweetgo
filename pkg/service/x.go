package service

import (
	"context"
	"os"

	xScraper "github.com/imperatrona/twitter-scraper"
)

type XService struct {
	scraper *xScraper.Scraper
}

func NewTwitterService() *XService {
	scraper := xScraper.New()
	scraper.SetAuthToken(xScraper.AuthToken{
		Token:     os.Getenv("TWITTER_TOKEN"),
		CSRFToken: os.Getenv("TWITTER_CSRF_TOKEN"),
	})

	if !scraper.IsLoggedIn() {
		panic("Invalid AuthToken")
	}
	return &XService{scraper: scraper}
}

func (x *XService) GetProfile(username string) (xScraper.Profile, error) {
	return x.scraper.GetProfile(username)
}

func (x *XService) GetTweets(ctx context.Context, username string, i int) <-chan *xScraper.TweetResult {
	return x.scraper.GetTweets(ctx, username, i)
}
