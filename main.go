package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var bearer string = "Bearer b778c82a75bf268bfece4a0e901c018c461f408b"

func main() {
	url := "https://www.strava.com/api/v3/athlete"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", bearer)
	fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string([]byte(body)))

}
