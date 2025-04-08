package models

import "time"

type ProfileResponse struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	TweetsCount int    `json:"tweets_count"`
}

type TweetResponse struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}
