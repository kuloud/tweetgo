# X Scraper API

A Go-based API for scraping X (formerly Twitter) profiles and tweets. This project provides endpoints to fetch user profiles and tweets using the `twitter-scraper` library.

## Features

- Fetch X user profiles by username.
- Fetch the latest tweets by username.
- Swagger API documentation for easy testing and integration.
- Rate limiting middleware for API protection.
- Environment variable support for sensitive credentials.

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Vercel CLI (for deployment)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/kuloud/tweetgo.git
   cd tweetgo

   ```

1. Install dependencies:

   ```bash
   go mod download

   ```

1. Set up environment variables: Create a `.env` file in the root directory with the following content:
   ```plainText
   TWITTER_TOKENS="token1:csrf1,token2:csrf2"
   ```

### API Documentation

Swagger UI is available at:

```plainText
http://{vercel.app.url}/swagger/
```

### Deployment to Vercel

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fkuloud%2Ftweetgo)

Steps

1. Install Vercel CLI:
   ```bash
   npm install -g vercel
   ```
1. Deploy the project:
   ```bash
   vercel
   ```
1. Set environment variables in the Vercel dashboard:

   - `TWITTER_TOKENS`: Twitter tokens in the format: TWITTER_TOKENS="token1:csrf1,token2:csrf2"
   - `ADMIN_USERNAME`: Your admin username (local development only)
   - `ADMIN_PASSWORD`: Your admin password (local development only)
   - `JWT_SECRET_KEY`: Your JWT secret key
   - `PORT`: Set to `8080` (local development only)

1. Access the deployed API at the provided Vercel URL.

### Access Information

- API Base URL: https://{vercel.app.url}/api/v1
- Swagger UI: https://{vercel.app.url}/swagger/

### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

### License

This project is governed by a Private License Agreement. Unauthorized use, reproduction, or distribution is strictly prohibited. For inquiries, please contact the project owner.
