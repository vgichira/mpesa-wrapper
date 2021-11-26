package mpesa

import (
	"bytes"
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

func (config *Config) LipaNaMpesaOnline(stkRequest *LipaNaMpesaRequest) (string, error) {
	requestBody, err := json.Marshal(stkRequest)

	if err != nil {
		return "", err
	}

	client := http.Client{}

	request, err := http.NewRequest("POST",
		fmt.Sprintf("%smpesa/stkpush/v1/processrequest", config.getBaseUrl()),
		bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	token, err := config.generateToken()

	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	return string(body), nil
}

func (config *Config) RegisterURL(requestBody *RegisterURL) (string, error) {
	body, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	client := http.Client{}

	request, err := http.NewRequest("POST",
		fmt.Sprintf("%smpesa/c2b/v1/registerurl",
			config.getBaseUrl()), bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	token, err := config.generateToken()

	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	apiResponse, err := ioutil.ReadAll(response.Body)

	return string(apiResponse), err
}
