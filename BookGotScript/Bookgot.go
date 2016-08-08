package Bookgot
/*
	BookGot is a simple boot for farm free books from PackPub.com, your functionality is based in
	verifying with multiple petitions if is available a new free book or no. This bot simulates be a client registered
	with your credentials, and if you haven't the last free book BookGot automatically claim the book for you.

	Author: Bregy Esteban Malpartida Ramos
	Date: 07 - 08 - 2016
	License: GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
*/

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"

	"net/http/cookiejar"
	"strings"

	"github.com/go-errors/errors"
	"io/ioutil"

	"golang.org/x/net/html"

	"github.com/yhat/scrape"

	"golang.org/x/net/html/atom"
	"net/url"
	"strconv"
	"time"
	
)

type PackPubUser struct {
	Email, Password string
	Logged          bool
}

const (
	PackPubUrl = "https://www.packtpub.com/"
)

func CreateNewPBUser(email, password string) *PackPubUser {
	runtime.GOMAXPROCS(4)
	var finalUser PackPubUser
	finalUser.Email = email
	finalUser.Password = password
	return &finalUser
}

func searchLinkInRawBody(page string) (string, error) {
	err1 := errors.New("Error en la respuesta, no encontrada la opcion de libro gratis")
	if strings.Contains(page, `<a href="/freelearning-claim/`) {
		if indexOfHref := strings.Index(page, `<a href="/freelearning-claim/`); indexOfHref != -1 {
			iFinal := indexOfHref + 40
			iInit := indexOfHref + 10
			hrefFreeBook := page[iInit:iFinal]
			fmt.Println(hrefFreeBook)

			return hrefFreeBook, nil

		} else {
			return "", err1
		}
	} else {
		return "", err1
	}
	return "", nil
}

func loginInPackPub(email, password string) (http.CookieJar, bool) {

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"email\"\r\n\r\n" + email + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"password\"\r\n\r\n" + password + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"op\"\r\n\r\nLogin\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"form_build_id\"\r\n\r\nform-028226298231fc14796523a08930e306\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"form_id\"\r\n\r\npackt_user_login_form\r\n-----011000010111000001101001--")

	req, _ := http.NewRequest("POST", PackPubUrl+"packt/offers/free-learning", payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("connection", "keep-alive")

	res, _ := client.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	//url, _ := url.Parse(PackPubUrl)
	cookies := client.Jar

	return cookies, true
}

func GetTimeForNewFreeBook() (time.Time, error) {
	req, _ := http.NewRequest("GET", PackPubUrl+"packt/offers/free-learning", nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	root, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return time.Time{}, err
	}
	timeTo, _ := scrape.Find(root, scrape.ByClass("packt-js-countdown"))
	sTimeTo := scrape.Attr(timeTo, "data-countdown-to")
	tm, _ := strconv.ParseInt(sTimeTo, 10, 64)
	t := time.Unix(tm, 0).UTC().Unix()
	actualTime := time.Now().UTC().Unix()

	countdown := t - actualTime
	finalCountdown := time.Unix(countdown, 0).UTC()

	return finalCountdown, nil
}

func verifyIfUserLogged(email, password string) bool {
	if cookies, r := loginInPackPub(email, password); r {
		url, _ := url.Parse(PackPubUrl)
		//fmt.Println(cookies)
		cookieSess := cookies.Cookies(url)[0]
		req, _ := http.NewRequest("GET", PackPubUrl+"packt/offers/free-learning", nil)
		req.AddCookie(cookieSess)
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		if strings.Contains(string(body), email) {
			return true
		} else {
			return false
		}

	}
	return false
}

func GetUrlOfFreeBook() (string, error) {
	req, err := http.NewRequest("GET", PackPubUrl+"packt/offers/free-learning", nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	fburl, err := searchLinkInRawBody(string(body))
	if err != nil {
		return "", err
	}
	return PackPubUrl + fburl, nil
}

func pullNewFreeBook(email, password string) bool {
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"email\"\r\n\r\n" + email + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"password\"\r\n\r\n" + password + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"op\"\r\n\r\nLogin\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"form_build_id\"\r\n\r\nform-028226298231fc14796523a08930e306\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"form_id\"\r\n\r\npackt_user_login_form\r\n-----011000010111000001101001--")

	req, _ := http.NewRequest("POST", PackPubUrl+"packt/offers/free-learning", payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("connection", "keep-alive")

	res, _ := client.Do(req)
	res.Body.Close()
	u, _ := url.Parse(PackPubUrl)
	fmt.Println(client.Jar.Cookies(u))
	fUrl, _ := GetUrlOfFreeBook()
	fmt.Println(fUrl)

	pushableBook, _ := http.NewRequest("GET", fUrl, nil)

	resp, _ := client.Do(pushableBook)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if strings.Contains(string(body), email) {
		return true
	} else {
		return false
	}
	return false
}

func GetNameOfCurrentFreeBook() string {
	req, _ := http.NewRequest("GET", PackPubUrl+"packt/offers/free-learning", nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	root, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "Not found"
	}
	classOfName, _ := scrape.Find(root, scrape.ByClass("dotd-title"))
	nameOfBook, _ := scrape.Find(classOfName, scrape.ByTag(atom.H2))
	nS := scrape.Text(nameOfBook)
	return nS
}

func getBooksFromUser(email, password string) []string {
	var wg sync.WaitGroup
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"email\"\r\n\r\n" + email + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"password\"\r\n\r\n" + password + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"op\"\r\n\r\nLogin\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"form_build_id\"\r\n\r\nform-028226298231fc14796523a08930e306\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"form_id\"\r\n\r\npackt_user_login_form\r\n-----011000010111000001101001--")

	req, _ := http.NewRequest("POST", PackPubUrl+"packt/offers/free-learning", payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("connection", "keep-alive")

	res, _ := client.Do(req)
	res.Body.Close()

	allBooks, _ := http.NewRequest("GET", "https://www.packtpub.com/account/my-ebooks", nil)

	resp, _ := client.Do(allBooks)

	defer resp.Body.Close()

	/*body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))*/

	z := html.NewTokenizer(resp.Body)
	booksOfUser := []string{}
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return booksOfUser
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "div"
			if isAnchor {
				for _, a := range t.Attr {
					wg.Add(1)
					go func(a html.Attribute, t html.Token, booksArray *[]string) {
						if a.Key == "class" {
							if a.Val == "product-line unseen" {
								title := t.Attr[2].Val
								*booksArray = append(*booksArray, title)
								//fmt.Println("libro: ", title)

							}
						}
						wg.Done()

					}(a, t, &booksOfUser)
				}

			}
		}
	}
	wg.Wait()
	return booksOfUser

}

func (user *PackPubUser) Login() bool {
	if verifyIfUserLogged(user.Email, user.Password) {
		user.Logged = true
		return true
	} else {
		user.Logged = false
		return false
	}
}

func (user *PackPubUser) GetNamesOfBooks() []string {
	return getBooksFromUser(user.Email, user.Password)
}

func (user *PackPubUser) VerifyIfIHaveLastFreeBook() bool {
	myBooks := user.GetNamesOfBooks()
	actualFreeBook := GetNameOfCurrentFreeBook()
	for _, book := range myBooks {

		comp := strings.Contains(book, actualFreeBook)
		if comp {
			fmt.Println(book, " is ", actualFreeBook)
			return true
		}

	}
	return false
}

func (user *PackPubUser) PullNewFreeBook() bool{
	if pullNewFreeBook(user.Email, user.Password){
		return true
	}else{
		return false
	}
}


