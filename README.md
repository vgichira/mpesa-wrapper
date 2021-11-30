# Safaricom Mpesa Golang SDK
This is an unofficial Mpesa Daraja API wrapper. You can easily install and reuse this package in you Go projects. This allows you to focus on the business logic by leaving the Mpesa APIs heavylifting.

## Installation

## Initialize instance of the gateway

```go
    mpesa, err := mpesa.Init("CONSUMER_KEY_HERE", "CONSUMER_SECRET_HERE", "ENVIROMENT (SANDBOX / LIVE)")
```