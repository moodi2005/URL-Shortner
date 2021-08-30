// save Link item in database
// chack unique some item

package database

import (
	"database/sql"
	"fmt"

	"github.com/amir-mhmd-najafi/URL-Shortner/urlshortener"
)

// save link data in databse
func SaveLinkInDatabase(linkData urlshortener.Link, DB *sql.DB) error {

	dbCommand := fmt.Sprintf(`INSERT INTO urlshortened (shortenedlink, shownumberofclicklink, notshortenedlink, numberofclick, ip, userid) VALUES ('%s','%s','%s', %d, '%s', %d);`,
		linkData.ShortenedLink, linkData.ShowNumberOfClickLink, linkData.NotShortenedLink, linkData.NumberOfClick, linkData.IP, linkData.UserID)
	_, err := DB.Exec(dbCommand)
	if err != nil {
		return err
	}

	return nil
}

// check link exists and return long link
func CheckExistsLink(want string, have string, value string, DB *sql.DB) (string, error) {
	dbCommand := fmt.Sprintf(`SELECT %s FROM urlshortened WHERE %s = '%s';`,
		want, have, value)
	var result string
	err := DB.QueryRow(dbCommand).Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}
