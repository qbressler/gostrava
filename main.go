package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/qbressler/stravaApp/configure"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getToken() (string, error) {
	//token url
	creds := configure.Get()
	tokenURL := "https://www.strava.com/oauth/token"
	data := url.Values{}
	data.Set("client_id", creds.ClientID)
	data.Set("client_secret", creds.ClientSecret)
	data.Set("code", creds.Code)

	client := &http.Client{}
	resp, err := client.PostForm(tokenURL, data)
	if err != nil {
		return "", err
	}

	var response AuthResponse
	respString, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respString, &response)
	fmt.Println(string(respString))
	return response.AccessToken, err
}

func main() {
	token, err := getToken()
	if err != nil {
		log.Fatal(err)
	}
	bearer := "Bearer " + token
	//url := "https://www.strava.com/api/v3/athlete"
	url := "https://www.strava.com/api/v3/activities/5190856387?include_all_efforts=true"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", bearer)
	fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string([]byte(body)))
}
