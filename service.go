package mpesa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
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

func (config *Config) generateLipaNaMpesaPassword(shortCode, passKey, timestamp string) string {
	return base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s%s%s", shortCode, passKey, timestamp)))
}

func (config *Config) LipaNaMpesaOnline(stkRequest *LipaNaMpesaRequest) (string, error) {
	requestBody, err := json.Marshal(stkRequest)

	if err != nil {
		return "", err
	}

	return config.makeRequest("POST", requestBody)
}

func (config *Config) RegisterURL(requestBody *RegisterURL) (string, error) {
	body, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	return config.makeRequest("POST", body)
}

func (config *Config) SimulateC2B(requestBody *C2BTransaction) (string, error) {
	body, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	return config.makeRequest("POST", body)
}

func (config *Config) InitiateB2C(requestBody *B2C) (string, error) {
	body, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	return config.makeRequest("POST", body)
}

func (config *Config) makeRequest(method string, body []byte) (string, error) {
	client := &http.Client{}

	request, err := http.NewRequest("POST",
		fmt.Sprintf("%smpesa/b2c/v1/paymentrequest",
			config.getBaseUrl()), bytes.NewReader(body))

	if err != nil {
		return "", err
	}

	// Generate token
	token, err := config.generateToken()

	if err != nil {
		return "", err
	}

	// Add Request Headers
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

func (config *Config) generateSecurityCredential(
	publicCertLocation, initiatorPassword string) (string, error) {
	var pubKey []byte

	resp, err := http.Get(publicCertLocation)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	pubKey, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(pubKey)

	var cert *x509.Certificate

	cert, _ = x509.ParseCertificate(block.Bytes)

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	rng := rand.Reader

	encrypted, err := rsa.EncryptPKCS1v15(rng, rsaPublicKey, []byte(initiatorPassword))

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
