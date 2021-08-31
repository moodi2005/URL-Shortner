package app

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/amir-mhmd-najafi/URL-Shortner/internal/database"
	"github.com/amir-mhmd-najafi/URL-Shortner/pkg/urlshortener"
	"github.com/lib/pq"
)

func Shortner(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	// check post methd
	// if other method => redirect to home
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	link, err := urlshortener.UrlShortener(w, r)
	// get link struct with data about link
	if err != nil {
		fmt.Println(err)
		return
	}

	// save unique in database
	// if not unique => genrate again random string
	for {
		if err = database.SaveLinkInDatabase(link, DB); err != nil {
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
	html := fmt.Sprintf(`http://localhost:5500/%s`, link.ShortenedLink)
	fmt.Fprint(w, html)

}
