package main

import (
	"net/http"

	"github.com/amir-mhmd-najafi/URL-Shortner/urlshortener"
)

func main() {
	http.HandleFunc("/", servehomePage)
	http.HandleFunc("/shortened", shortened)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func servehomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../UI/Html/index.html")
}

func shortened(w http.ResponseWriter, r *http.Request) {
	urlshortener.UrlShortener(w, r)
}
