package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/amir-mhmd-najafi/URL-Shortner/database"
	"github.com/amir-mhmd-najafi/URL-Shortner/database/databaseconfig"
	"github.com/amir-mhmd-najafi/URL-Shortner/urlshortener"
	"github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error // test!
	DB, err = databaseconfig.ConnectToDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/shortened", shortened)
	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		panic(err)
	}
}

// home page => input for link => shortened link
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../UI/Html/index.html")
}

func shortened(w http.ResponseWriter, r *http.Request) {

	// get link struct with data about link
	link, err := urlshortener.UrlShortener(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = database.SaveSpecifiedLinkInDatabase(link, DB); err != nil {
		fmt.Println(err)
		return
	}

	// save random link data in database
	// if not unique, call GenerateRandomLinkAgain function for edit random link
	for {
		if err = database.SaveRandomLinkInDatabase(link, DB); err != nil {
			pqErrorCode := err.(*pq.Error).Code
			if pqErrorCode == "23505" {
				link = urlshortener.GenerateRandomLinkAgain()
				continue
			}
			fmt.Println(err)
			return
		}
		break
	}

}
