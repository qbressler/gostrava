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

// authentication
// get code
//https://www.strava.com/oauth/authorize?client_id=18876&response_type=code&redirect_uri=http://localhost/exchange_token&approval_prompt=force&scope=read,activity:read

// get bearer token:
//curl -X POST https://www.strava.com/oauth/token -F client_id=18876 -F client_secret=0d43a7363c7d2ab93ee6f34f2e2a64b11ed38cb8 -F code=0c3117d    18f7ca3b4a1fdbe9c27ee8b0adb8c41ed

//var bearer string = "Bearer cb5f8fad8da1290fb6e36098f4f001d1b51978db"

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
