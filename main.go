package main

import (
	"fmt"
	"log"
	"net/http"
)

var urlStore = map[string]string{}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/shorten", ShortenHandler)
	http.HandleFunc("/r/", RedirectHandler)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to URL Shortener!")
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	longURL := r.FormValue("url")
	shortCode := fmt.Sprintf("%x", len(urlStore)+1)
	urlStore[shortCode] = longURL
	fmt.Fprintf(w, "Short URL: http://localhost:8080/r/%s\n", shortCode)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/r/"):]
	longURL, ok := urlStore[code]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
