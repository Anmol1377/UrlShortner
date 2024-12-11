package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"Creation_data"`
}

// sample data
/* d87dyd8d ----> {
	ID = "d87dyd8d",
	OriginalUrl : "https://inc42.com",
	ShortURL : "d87dyd8d",
	CreationDate : "time.now()"
 }
*/

var urlDB = make(map[string]URL)

func genrateShortUrl(OriginalUrl string) string {

	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	data := hasher.Sum(nil)
	// fmt.Println("hasher:" , data)
	hash := hex.EncodeToString(data)
	fmt.Println(hash[:8])
	return hash[:8]
}

func createURL(OriginalUrl string) string {
	shortURL := genrateShortUrl(OriginalUrl)
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

func getURL(id string) (URL, error) {

	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("url not found")
	}
	return url, nil

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Method")
	fmt.Fprintf(w, "hello world")

}

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

	// fmt.Fprintf(w, shortUrl)

	response := struct {
		ShortedURL  string `json:"short_url"`
		OriginalUrl string `json:"original_url"`
	}{ShortedURL: shortUrl, OriginalUrl: data.URL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "invalid request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)

}

func main() {
	fmt.Println("Starting URL shoortner")
	OriginalUrl := "https://inc42.com"
	genrateShortUrl(OriginalUrl)

	//server

	http.HandleFunc("/", handler)
	http.HandleFunc("/shortn", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	fmt.Println("starting")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("starting issue", err)
	}

}
