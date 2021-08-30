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

	dbCommand := fmt.Sprintf(`INSERT INTO urlshortened (ShortenedLink, ShowNumberOfClickLink, NotShortenedLink, NumberOfClick, IP, UserID) VALUES ('%s','%s','%s', %d, '%s', %d);`,
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

// add 1 in link click count
func UpdateLinkCount(link string, DB *sql.DB) error {
	dbCommand := fmt.Sprintf(`UPDATE urlshortened SET NumberOfClick = NumberOfClick + 1 WHERE ShortenedLink = '%s';`, 
link)
_, err := DB.Exec(dbCommand)
	if err != nil {
		return err
	}

	return nil
}