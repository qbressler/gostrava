package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var pass string = getPass()
var authUrl string = "http://www.strava.com/oauth/authorize?client_id=18876&response_type=code&redirect_uri=http://localhost:8080/exchangeToken?approval_prompt=force&scope=read,activity:read"
var bearerToken string

type AuthResponse struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/exchangeToken", exchangeToken)
	http.HandleFunc("/app", app)
	http.ListenAndServe(":8080", nil)

}

func getPass() string {
	fbytes, err := ioutil.ReadFile("settings/thepass.txt")
	if err != nil {
		log.Print(err)
	}
	log.Println(string(fbytes))
	return string(fbytes)
}

func app(w http.ResponseWriter, r *http.Request) {
	if bearerToken == "" {
		fmt.Fprint(w, "<h2>You are not authenticated anymore...</h2>")
	} else {

		fmt.Fprintf(w, `
			<h1>Welcome to the app</h1>
			You have a bearer token...and that token is... %s<br />
			What would you like to do now?

		`, bearerToken)
	}
}
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home page... give all your access to us!</h1>")
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)

}

func exchangeToken(w http.ResponseWriter, r *http.Request) {
	var code string
	splt := strings.Split(r.URL.RawQuery, "&")
	for _, j := range splt {
		t := j[:4]
		if t == "code" {
			s := strings.Split(j, "=")
			code = s[1]
		}
	}

	v := url.Values{}
	v.Set("client_id", "18876")
	v.Set("client_secret", pass)
	v.Set("code", code)
	v.Set("grant_type", "authorization_code")

	client := &http.Client{}
	log.Print(strings.NewReader(v.Encode()))
	r, err := http.NewRequest("POST", "https://www.strava.com/oauth/token", strings.NewReader(v.Encode()))
	if err != nil {
		log.Print(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	// build authResponse struct{}
	var authResponse AuthResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		log.Print(err)
	}
	log.Print(authResponse)

	bearerToken = authResponse.Token

	http.Redirect(w, r, "/app", http.StatusTemporaryRedirect)

}
