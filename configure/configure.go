package configure

import (
	"encoding/json"
	"log"
	"os"
)

type AuthCodes struct {
	ClientID     string
	ClientSecret string
	Code         string
}

func Get() AuthCodes {

	jsonSettingsFile, err := os.ReadFile("client_settings.json")
	if err != nil {
		log.Println(err)
	}

	var authCodes AuthCodes

	json.Unmarshal(jsonSettingsFile, &authCodes)
	return AuthCodes{
		ClientID:     "18876",
		ClientSecret: "0d43a7363c7d2ab93ee6f34f2e2a64b11ed38cb8",
		Code:         "d7585a9463a2b07958e92962615e2c38e662dae0",
	}
}
