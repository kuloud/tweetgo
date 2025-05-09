package service

import (
	"context"
	"sync"

	xScraper "github.com/imperatrona/twitter-scraper"
)

type XService struct {
	scraper *xScraper.Scraper
}

type XServicePool struct {
	services []*XService
	mu       sync.Mutex
	index    int
}

func NewXServicePool(tokens []xScraper.AuthToken) *XServicePool {
	pool := &XServicePool{}
	for _, token := range tokens {
		scraper := xScraper.New()
		scraper.SetAuthToken(token)
		if !scraper.IsLoggedIn() {
			panic("Invalid AuthToken")
		}
		pool.services = append(pool.services, &XService{scraper: scraper})
	}
	return pool
}

func (p *XServicePool) getNextService() *XService {
	p.mu.Lock()
	defer p.mu.Unlock()
	service := p.services[p.index]
	p.index = (p.index + 1) % len(p.services)
	return service
}

func (p *XServicePool) GetProfile(username string) (xScraper.Profile, error) {
	service := p.getNextService()
	return service.scraper.GetProfile(username)
}

func (p *XServicePool) GetTweets(ctx context.Context, username string, i int) <-chan *xScraper.TweetResult {
	service := p.getNextService()
	return service.scraper.GetTweets(ctx, username, i)
}
func (p *XServicePool) GetTweet(tweetId string) (*xScraper.Tweet, error) {
	service := p.getNextService()
	return service.scraper.GetTweet(tweetId)
}
