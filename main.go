package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/exchangeToken", exchangeToken)
	http.ListenAndServe(":8080", nil)

}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home page... give all your access to us!</h1>")
	url := "http://www.strava.com/oauth/authorize?client_id=18876&response_type=code&redirect_uri=http://localhost:8080/exchangeToken?approval_prompt=force&scope=read,activity:read" //todo get URL!

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func exchangeToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are exchanging tokens...good job :)")
	//	http://localhost:8080/exchangeToken?approval_prompt=force&state=&code=c597abeaba0b5526339e96cec26cd4b815e2eae9&scope=read,activity:read
	splt := strings.Split(r.URL.RawQuery, "&")
	for _, j := range splt {
		t := j[:4]
		if t == "code" {
			s := strings.Split(j, "=")
			fmt.Fprint(w, s[1])
		}
	}
	/*
		POST to oauth and get bearer token using the code you received from above!
		curl -X POST https://www.strava.com/oauth/token \
		-F client_id=YOURCLIENTID \
		-F client_secret=YOURCLIENTSECRET \
		-F code=AUTHORIZATIONCODE \
		-F grant_type=authorization_code
	*/
}
