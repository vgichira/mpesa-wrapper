package mpesa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Init(consumerKey, consumerSecret, environment string) *Config {
	return &Config{consumerKey, consumerSecret, environment}
}

func (config *Config) getBaseUrl() string {
	baseUrl := "https://sandbox.safaricom.co.ke/"

	if config.Environment == "PRODUCTION" {
		baseUrl = "https://api.safaricom.co.ke/"
	}

	return baseUrl
}

func (config *Config) generateToken() (string, error) {
	// Get the API Base Url based on the environment
	// PRODUCTION / SANDBOX
	baseUrl := config.getBaseUrl()

	endpoint := fmt.Sprintf("%soauth/v1/generate?grant_type=client_credentials", baseUrl)

	// Create request
	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return "", err
	}

	request.SetBasicAuth(config.ConsumerKey, config.ConsumerSecret)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	var apiResponse map[string]string

	err = json.Unmarshal(body, &apiResponse)

	if err != nil {
		return "", err
	}

	token := apiResponse["access_token"]

	return token, nil
}
