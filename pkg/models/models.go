package models

import "time"

// Profile of twitter user.
type Profile struct {
	Avatar               string     `json:"Avatar"`
	Banner               string     `json:"Banner"`
	Biography            string     `json:"Biography"`
	Birthday             string     `json:"Birthday"`
	FollowersCount       int        `json:"FollowersCount"`
	FollowingCount       int        `json:"FollowingCount"`
	FriendsCount         int        `json:"FriendsCount"`
	IsPrivate            bool       `json:"IsPrivate"`
	IsVerified           bool       `json:"IsVerified"`
	IsBlueVerified       bool       `json:"IsBlueVerified"`
	Joined               *time.Time `json:"Joined"`
	LikesCount           int        `json:"LikesCount"`
	ListedCount          int        `json:"ListedCount"`
	Location             string     `json:"Location"`
	Name                 string     `json:"Name"`
	PinnedTweetIDs       []string   `json:"PinnedTweetIDs"`
	TweetsCount          int        `json:"TweetsCount"`
	URL                  string     `json:"URL"`
	UserID               string     `json:"UserID"`
	Username             string     `json:"Username"`
	Website              string     `json:"Website"`
	Sensitive            bool       `json:"Sensitive"`
	Following            bool       `json:"Following"`
	FollowedBy           bool       `json:"FollowedBy"`
	MediaCount           int        `json:"MediaCount"`
	FastFollowersCount   int        `json:"FastFollowersCount"`
	NormalFollowersCount int        `json:"NormalFollowersCount"`
	ProfileImageShape    string     `json:"ProfileImageShape"`
	HasGraduatedAccess   bool       `json:"HasGraduatedAccess"`
	CanHighlightTweets   bool       `json:"CanHighlightTweets"`
}

type (
	// Mention type.
	Mention struct {
		ID       string `json:"ID"`
		Username string `json:"Username"`
		Name     string `json:"Name"`
	}

	// Url represents a URL with display, expanded, and index data.
	Url struct {
		DisplayURL  string `json:"display_url"`
		ExpandedURL string `json:"expanded_url"`
		URL         string `json:"url"`
		Indices     []int  `json:"indices"`
	}

	// Photo type.
	Photo struct {
		ID  string `json:"ID"`
		URL string `json:"URL"`
	}

	// Video type.
	Video struct {
		ID      string `json:"ID"`
		Preview string `json:"Preview"`
		URL     string `json:"URL"`
		HLSURL  string `json:"HLSURL"`
	}

	// GIF type.
	GIF struct {
		ID      string `json:"ID"`
		Preview string `json:"Preview"`
		URL     string `json:"URL"`
	}

	// Tweet type.
	Tweet struct {
		ConversationID    string    `json:"ConversationID"`
		GIFs              []GIF     `json:"GIFs"`
		Hashtags          []string  `json:"Hashtags"`
		HTML              string    `json:"HTML"`
		ID                string    `json:"ID"`
		InReplyToStatus   *Tweet    `json:"InReplyToStatus"`
		InReplyToStatusID string    `json:"InReplyToStatusID"`
		IsQuoted          bool      `json:"IsQuoted"`
		IsPin             bool      `json:"IsPin"`
		IsReply           bool      `json:"IsReply"`
		IsRetweet         bool      `json:"IsRetweet"`
		IsSelfThread      bool      `json:"IsSelfThread"`
		Likes             int       `json:"Likes"`
		Name              string    `json:"Name"`
		Mentions          []Mention `json:"Mentions"`
		PermanentURL      string    `json:"PermanentURL"`
		Photos            []Photo   `json:"Photos"`
		Place             *Place    `json:"Place"`
		QuotedStatus      *Tweet    `json:"QuotedStatus"`
		QuotedStatusID    string    `json:"QuotedStatusID"`
		Replies           int       `json:"Replies"`
		Retweets          int       `json:"Retweets"`
		RetweetedStatus   *Tweet    `json:"RetweetedStatus"`
		RetweetedStatusID string    `json:"RetweetedStatusID"`
		Text              string    `json:"Text"`
		Thread            []*Tweet  `json:"Thread"`
		TimeParsed        time.Time `json:"TimeParsed"`
		Timestamp         int64     `json:"Timestamp"`
		URLs              []string  `json:"URLs"`
		UserID            string    `json:"UserID"`
		Username          string    `json:"Username"`
		Videos            []Video   `json:"Videos"`
		Views             int       `json:"Views"`
		SensitiveContent  bool      `json:"SensitiveContent"`
	}

	// ProfileResult of scrapping.
	ProfileResult struct {
		Profile `json:"Profile"`
		Error   error `json:"Error"`
	}

	// TweetResult of scrapping.
	TweetResult struct {
		Tweet `json:"Tweet"`
		Error error `json:"Error"`
	}

	ScheduledTweet struct {
		ID        string    `json:"ID"`
		State     string    `json:"State"`
		ExecuteAt time.Time `json:"ExecuteAt"`
		Text      string    `json:"Text"`
		Videos    []Video   `json:"Videos"`
		Photos    []Photo   `json:"Photos"`
		GIFs      []GIF     `json:"GIFs"`
	}

	ExtendedMedia struct {
		IDStr                    string `json:"id_str"`
		MediaURLHttps            string `json:"media_url_https"`
		ExtSensitiveMediaWarning struct {
			AdultContent    bool `json:"adult_content"`
			GraphicViolence bool `json:"graphic_violence"`
			Other           bool `json:"other"`
		} `json:"ext_sensitive_media_warning"`
		Type      string `json:"type"`
		URL       string `json:"url"`
		VideoInfo struct {
			Variants []struct {
				Type    string `json:"content_type"`
				Bitrate int    `json:"bitrate"`
				URL     string `json:"url"`
			} `json:"variants"`
		} `json:"video_info"`
	}

	legacyTweet struct {
		ConversationIDStr string `json:"conversation_id_str"`
		CreatedAt         string `json:"created_at"`
		FavoriteCount     int    `json:"favorite_count"`
		FullText          string `json:"full_text"`
		Entities          struct {
			Hashtags []struct {
				Text string `json:"text"`
			} `json:"hashtags"`
			Media []struct {
				MediaURLHttps string `json:"media_url_https"`
				Type          string `json:"type"`
				URL           string `json:"url"`
			} `json:"media"`
			URLs         []Url `json:"urls"`
			UserMentions []struct {
				IDStr      string `json:"id_str"`
				Name       string `json:"name"`
				ScreenName string `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []ExtendedMedia `json:"media"`
		} `json:"extended_entities"`
		IDStr                 string `json:"id_str"`
		InReplyToStatusIDStr  string `json:"in_reply_to_status_id_str"`
		Place                 Place  `json:"place"`
		ReplyCount            int    `json:"reply_count"`
		RetweetCount          int    `json:"retweet_count"`
		RetweetedStatusIDStr  string `json:"retweeted_status_id_str"`
		RetweetedStatusResult struct {
			Result *result `json:"result"`
		} `json:"retweeted_status_result"`
		QuotedStatusIDStr string `json:"quoted_status_id_str"`
		SelfThread        struct {
			IDStr string `json:"id_str"`
		} `json:"self_thread"`
		Time      time.Time `json:"time"`
		UserIDStr string    `json:"user_id_str"`
		Views     struct {
			State string `json:"state"`
			Count string `json:"count"`
		} `json:"ext_views"`
	}

	legacyUser struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		Entities    struct {
			URL struct {
				Urls []struct {
					ExpandedURL string `json:"expanded_url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount         int      `json:"favourites_count"`
		FollowersCount          int      `json:"followers_count"`
		FriendsCount            int      `json:"friends_count"`
		IDStr                   string   `json:"id_str"`
		ListedCount             int      `json:"listed_count"`
		Name                    string   `json:"name"`
		Location                string   `json:"location"`
		PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL        string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string   `json:"profile_image_url_https"`
		Protected               bool     `json:"protected"`
		ScreenName              string   `json:"screen_name"`
		StatusesCount           int      `json:"statuses_count"`
		Verified                bool     `json:"verified"`
		FollowedBy              bool     `json:"followed_by"`
		Following               bool     `json:"following"`
		CanDm                   bool     `json:"can_dm"`
		CanMediaTag             bool     `json:"can_media_tag"`
		DefaultProfile          bool     `json:"default_profile"`
		DefaultProfileImage     bool     `json:"default_profile_image"`
		FastFollowersCount      int      `json:"fast_followers_count"`
		HasCustomTimelines      bool     `json:"has_custom_timelines"`
		IsTranslator            bool     `json:"is_translator"`
		MediaCount              int      `json:"media_count"`
		NeedsPhoneVerification  bool     `json:"needs_phone_verification"`
		NormalFollowersCount    int      `json:"normal_followers_count"`
		PossiblySensitive       bool     `json:"possibly_sensitive"`
		ProfileInterstitialType string   `json:"profile_interstitial_type"`
		TranslatorType          string   `json:"translator_type"`
		WantRetweets            bool     `json:"want_retweets"`
		WithheldInCountries     []string `json:"withheld_in_countries"`
	}

	legacyUserV2 struct {
		FollowedBy          bool   `json:"followed_by"`
		Following           bool   `json:"following"`
		CanDm               bool   `json:"can_dm"`
		CanMediaTag         bool   `json:"can_media_tag"`
		CreatedAt           string `json:"created_at"`
		DefaultProfile      bool   `json:"default_profile"`
		DefaultProfileImage bool   `json:"default_profile_image"`
		Description         string `json:"description"`
		Entities            struct {
			Description struct {
				Urls []Url `json:"urls"`
			} `json:"description"`
			URL struct {
				Urls []Url `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FastFollowersCount      int           `json:"fast_followers_count"`
		FavouritesCount         int           `json:"favourites_count"`
		FollowersCount          int           `json:"followers_count"`
		FriendsCount            int           `json:"friends_count"`
		HasCustomTimelines      bool          `json:"has_custom_timelines"`
		IsTranslator            bool          `json:"is_translator"`
		ListedCount             int           `json:"listed_count"`
		Location                string        `json:"location"`
		MediaCount              int           `json:"media_count"`
		Name                    string        `json:"name"`
		NormalFollowersCount    int           `json:"normal_followers_count"`
		PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
		PossiblySensitive       bool          `json:"possibly_sensitive"`
		ProfileBannerURL        string        `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
		ProfileInterstitialType string        `json:"profile_interstitial_type"`
		ScreenName              string        `json:"screen_name"`
		StatusesCount           int           `json:"statuses_count"`
		TranslatorType          string        `json:"translator_type"`
		URL                     string        `json:"url"`
		Verified                bool          `json:"verified"`
		WantRetweets            bool          `json:"want_retweets"`
		WithheldInCountries     []interface{} `json:"withheld_in_countries"`
	}

	Place struct {
		ID          string `json:"id"`
		PlaceType   string `json:"place_type"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		CountryCode string `json:"country_code"`
		Country     string `json:"country"`
		BoundingBox struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"bounding_box"`
	}

	fetchProfileFunc func(query string, maxProfilesNbr int, cursor string) ([]*Profile, string, error)
	fetchTweetFunc   func(query string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)

	legacyExtendedProfile struct {
		Birthdate struct {
			Day            int    `json:"day"`
			Month          int    `json:"month"`
			Year           int    `json:"year"`
			Visibility     string `json:"visibility"`
			YearVisibility string `json:"year_visibility"`
		} `json:"birthdate"`
	}

	verificationInfo struct {
		IsIdentityVerified bool `json:"is_identity_verified"`
		Reason             struct {
			Description struct {
				Text     string `json:"text"`
				Entities []struct {
					FromIndex int `json:"from_index"`
					ToIndex   int `json:"to_index"`
					Ref       struct {
						URL     string `json:"url"`
						URLType string `json:"url_type"`
					} `json:"ref"`
				} `json:"entities"`
			} `json:"description"`
			VerifiedSinceMsec string `json:"verified_since_msec"`
		} `json:"reason"`
	}

	highlightsInfo struct {
		CanHighlightTweets bool   `json:"can_highlight_tweets"`
		HighlightedTweets  string `json:"highlighted_tweets"`
	}
)
