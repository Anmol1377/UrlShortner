# URL Shortener in Go

This is a simple URL shortener service built using Go. It allows you to shorten URLs and redirect to the original URL using a short link. The project uses a simple hash function (MD5) to generate unique short URLs for each original URL.

## Features

- **Generate Short URLs**: Accepts an original URL and returns a shortened version.
- **Redirect**: Redirects users from the shortened URL to the original URL.
- **Store Data**: The original URL and its short version are stored in memory.

## Technologies Used

- **Go (Golang)**: The backend language for implementing the URL shortener.
- **HTTP Server**: The project uses the `net/http` package to create an HTTP server for serving requests.
- **MD5 Hashing**: To generate a unique 8-character short URL based on the original URL.

## Project Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/url-shortener.git
    cd url-shortener
    ```

2. **Install Go** (if you don't have it installed):
    Follow the instructions at https://golang.org/doc/install to install Go.

3. **Run the application**:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:3000`.

## Endpoints

### 1. Create Short URL

**URL**: `/shortn`

**Method**: `POST`

**Request Body**:
```json
{
  "url": "https://example.com"
}
 ```
 **Response Body**:
    ```json
    {
      "short_url": "http://localhost:3000/redirect/d87dyd8d",
      "original_url": "https://example.com"
    }
    ```

