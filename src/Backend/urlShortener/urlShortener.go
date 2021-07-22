// ckage for genrate shortened link and show number of link clicked
// and saven Link struct
// then save to postgresql

package urlshortener

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Link struct {
	NotShortenedLink      string
	ShortenedLink         string
	NumberOfClick         uint
	ShowNumberOfClickLink string
}

var charset []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

// get not shortened link from input name=="notshortenedlink"
func getLink(w http.ResponseWriter, r *http.Request, link *Link) error {
	err := r.ParseForm()
	if err != nil {
		return fmt.Errorf("error in get not shortened link from input: %s", err)
	}
	link.NotShortenedLink = r.PostFormValue("notshortenedlink")
	return nil
}

func genrateShortenedLink(link *Link) {

	for {

		// create random string
		randomString := makeRandomString(4)

		// if exists in db, repeat the loop
		// else  save in link.ShortenedLink
		if checkLinkExists(randomString) {
			link.ShortenedLink = randomString
			break
		}

	}

}

// check db for exists link string
func checkLinkExists(link string) bool {
	// TODO
	_ = link
	return true
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
	if err := getLink(w, r, &link); err != nil {
		fmt.Println(err)
	}

	// 2. gnrate random string for shortened link and check not exists
	// and save in link.ShortenedLink
	genrateShortenedLink(&link)


	// 3. genrate ShowNumberOfClickLink and save to link.ShowNumberOfClickLink
	link.ShowNumberOfClickLink = makeRandomString(4)

	fmt.Println(link)
}

