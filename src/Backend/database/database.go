// save Link item in database
// chack unique some item

package database

import (
	"database/sql"
	"fmt"

	"github.com/amir-mhmd-najafi/URL-Shortner/urlshortener"
)

// save specified link data in databse
func SaveSpecifiedLinkInDatabase(linkData urlshortener.Link, DB *sql.DB) error {

	dbCommand := fmt.Sprintf(`INSERT INTO urlshortened (notshortenedlink, numberofclick, ip, userid) VALUES ('%s', %d, '%s', %d);`,
		linkData.NotShortenedLink, linkData.NumberOfClick, linkData.IP, linkData.UserID)
	_, err := DB.Exec(dbCommand)
	if err != nil {
		return err
	}

	return nil
}

// save random link data in databse
// if not shortened link unique, not saving
func SaveRandomLinkInDatabase(linkData urlshortener.Link, DB *sql.DB) error {

	dbCommand := fmt.Sprintf(`INSERT INTO urlshortened (shortenedlink, shownumberofclicklink) VALUES ('%s', '%s');`,
		linkData.ShortenedLink, linkData.ShowNumberOfClickLink)
	_, err := DB.Exec(dbCommand)
	if err != nil {
		return err
	}

	return nil
}