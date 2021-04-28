package configure

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AuthCodes struct {
	ClientID     string
	ClientSecret string
	Code         string
}

func Get() AuthCodes {

	jsonSettingsFile, err := ioutil.ReadFile("./client_settings.json")
	if err != nil {
		log.Println(err)
	}

	var authCodes AuthCodes

	json.Unmarshal(jsonSettingsFile, &authCodes)
	return authCodes
}
