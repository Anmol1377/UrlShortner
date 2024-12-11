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

## Code Explanation

### URL Struct

The `URL` struct represents a URL object and contains the following fields:

```go
type URL struct {
    ID           string    `json:"id"`           // Unique identifier (hashed from the original URL)
    OriginalUrl  string    `json:"original_url"` // Original, long URL
    ShortURL     string    `json:"short_url"`    // Shortened URL
    CreationDate time.Time `json:"creation_date"` // Timestamp when the URL was shortened
}
```



### generateShortUrl Function
This function generates a short URL by applying an MD5 hash to the original URL and extracting the first 8 characters of the hash.
```go
func generateShortUrl(OriginalUrl string) string {
    hasher := md5.New()
    hasher.Write([]byte(OriginalUrl))
    data := hasher.Sum(nil)
    hash := hex.EncodeToString(data)
    return hash[:8]
}
```

### createURL Function
This function generates a short URL, stores it in an in-memory database (urlDB), and returns the full URL with the base path (http://localhost:3000/redirect/{short_url}).


```go
func createURL(OriginalUrl string) string {
    shortURL := generateShortUrl(OriginalUrl)
    id := shortURL
    urlDB[id] = URL{
        ID:           id,
        OriginalUrl:  OriginalUrl,
        ShortURL:     shortURL,
        CreationDate: time.Now(),
    }
    baseURL := "http://localhost:3000"
    return fmt.Sprintf("%s/redirect/%s", baseURL, shortURL)
}

```
### getURL Function
This function retrieves the stored URL by its ID (the shortened URL's ID).

```go
func getURL(id string) (URL, error) {
    url, ok := urlDB[id]
    if !ok {
        return URL{}, errors.New("url not found")
    }
    return url, nil
}
```
### ShortURLHandler Function
This HTTP handler processes the POST request for creating a short URL. It decodes the JSON body, generates a short URL, and returns the result.


```go
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
    var data struct {
        URL string `json:"url"`
    }
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "invalid", http.StatusBadRequest)
        return
    }
    shortUrl := createURL(data.URL)
    response := struct {
        ShortedURL  string `json:"short_url"`
        OriginalUrl string `json:"original_url"`
    }{ShortedURL: shortUrl, OriginalUrl: data.URL}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

```


### redirectURLHandler Function
This handler performs the redirect from the short URL to the original URL based on the short URL's ID.


```go
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/redirect/"):]
    url, err := getURL(id)
    if err != nil {
        http.Error(w, "invalid request", http.StatusNotFound)
        return
    }
    http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

```

### Main Function
The main function initializes the server and sets up the necessary routes for handling requests.


```go
func main() {
    fmt.Println("Starting URL shortener")
    OriginalUrl := "https://inc42.com"
    generateShortUrl(OriginalUrl)

    // Server setup
    http.HandleFunc("/", handler)
    http.HandleFunc("/shortn", ShortURLHandler)
    http.HandleFunc("/redirect/", redirectURLHandler)
    fmt.Println("Starting server...")
    err := http.ListenAndServe(":3000", nil)

    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

```





