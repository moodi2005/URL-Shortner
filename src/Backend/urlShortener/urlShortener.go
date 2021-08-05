// ckage for generate shortened link and show number of link clicke
// and saven Link struct
// then save to postgresql

package urlshortener

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Link struct {
	ShortenedLink         string
	NotShortenedLink      string
	NumberOfClick         uint
	ShowNumberOfClickLink string
	IP                    string
	// TODO get user information if logined
}

var charset []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

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

	// generate ShowNumberOfClickLink
	for {
		// create random string
		randomString := makeRandomString(4)

		// if exists in db, repeat the loop
		// else  save in link.ShortenedLink
		if checkLinkExists(randomString, "ShortenedLink") {
			link.ShortenedLink = randomString
			break
		}

	}

	// generate ShowNumberOfClickLink
	for {
		// create random string
		randomString := makeRandomString(6)

		// if exists in db, repeat the loop
		// else  save in link.ShowNumberOfClickLink
		if checkLinkExists(randomString, "ShowNumberOfClickLink") {
			link.ShowNumberOfClickLink = randomString
			break
		}

	}

}

// check db for exists link string
func checkLinkExists(link string, field string) bool {
	// TODO
	_ = link
	_ = field
	return true
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
	randomString := make([]rune, randomStringLen)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func UrlShortener(w http.ResponseWriter, r *http.Request) {
	var link Link

	// 1. get and save link.NotShortenedLink
	if err := getNotShortenedLink(w, r, &link); err != nil {
		fmt.Println(err)
	}

	// 2. gnrate random string for shortened link and check not exists
	// and save in link.ShortenedLink
	// 3. generate ShowNumberOfClickLink and save to link.ShowNumberOfClickLink
	generateRandomLink(&link)

	// 4.get user IP
	getuserIP(r, &link)

	fmt.Println(link)
}
