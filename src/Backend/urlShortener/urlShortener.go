// ckage for generate shortened link and show number of link clicke
// and saven Link struct
// then save to postgresql

package urlshortener

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Link struct {
	ShortenedLink         string
	NotShortenedLink      string
	NumberOfClick         uint
	ShowNumberOfClickLink string
	IP                    string
	UserID                int
}

// link data save in this var
var link Link

// for create random string
const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

// for create random string
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// get not shortened link from input name=="notshortenedlink"
func getNotShortenedLink(w http.ResponseWriter, r *http.Request, link *Link) error {
	err := r.ParseForm()
	if err != nil {
		return fmt.Errorf("error in get not shortened link from input: %s", err)
	}
	link.NotShortenedLink = r.PostFormValue("notshortenedlink")
	return nil
}

func generateRandomLink(link *Link) {

	// generate ShowNumberOfClickLink, ShortenedLink
	link.ShortenedLink = makeRandomString(4)
	link.ShowNumberOfClickLink = makeRandomString(6)

}

// get ip from header
func getuserIP(r *http.Request, link *Link) {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		link.IP = forwarded
	}
	link.IP = r.RemoteAddr
}

// return randon string
func makeRandomString(randomStringLen int) string {
	randomString := make([]byte, randomStringLen)
	for i := range randomString {
		randomString[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomString)
}

// save link items by call other function
// and return for save in database in other package
func UrlShortener(w http.ResponseWriter, r *http.Request) (Link, error) {

	// 1. get and save link.NotShortenedLink
	if err := getNotShortenedLink(w, r, &link); err != nil {
		return Link{}, fmt.Errorf("error in get shortened link from input: %s", err)
	}
	// 2.get user IP
	getuserIP(r, &link)

	// 3. get user id
	//TODO

	// 4. genrate random string for shortened link
	// and save in link.ShortenedLink
	// 5. generate ShowNumberOfClickLink and save to link.ShowNumberOfClickLink
	generateRandomLink(&link)

	return link, nil
}

// genrate random string for shortened link and show number of click link in it is not unique
func GenerateRandomLinkAgain() Link {
	// 4. genrate random string for shortened link again
	// and save in link.ShortenedLink
	// 35. generate ShowNumberOfClickLink and save to link.ShowNumberOfClickLink again
	generateRandomLink(&link)

	return link
}
