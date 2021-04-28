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

var pass string = "0d43a7363c7d2ab93ee6f34f2e2a64b11ed38cb8"

type AuthResponse struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/exchangeToken", exchangeToken)
	http.ListenAndServe(":8080", nil)

}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home page... give all your access to us!</h1>")
	url := "http://www.strava.com/oauth/authorize?client_id=18876&response_type=code&redirect_uri=http://192.168.1.147:8080/exchangeToken?approval_prompt=force&scope=read,activity:read" //todo get URL!

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

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
	var authResponse AuthResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		log.Print(err)
	}
	log.Print(authResponse)

}
