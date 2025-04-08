module handler

go 1.24.2

require (
	github.com/gorilla/mux v1.8.1
	github.com/imperatrona/twitter-scraper v0.0.17
	github.com/joho/godotenv v1.5.1
	github.com/swaggo/http-swagger v1.3.4
	github.com/swaggo/swag v1.16.4
)

require (
	github.com/AlexEidt/Vidio v1.5.1 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/aws/aws-lambda-go v1.47.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/swaggo/files v1.0.1 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/kuloud/tweetgo v0.0.0-unpublished

replace github.com/kuloud/tweetgo => ./

require github.com/vercel/go-bridge v0.0.0-20221108222652-296f4c6bdb6d
