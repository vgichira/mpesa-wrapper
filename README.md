# Safaricom Mpesa Golang SDK
This is an unofficial Mpesa Daraja API wrapper. You can easily install and reuse this package in you Go projects. This allows you to focus on the business logic by leaving the Mpesa APIs heavylifting.

## Installation
This documentation assumes that you already have Go installed on your device. If you have not installed Go, check out the [documentation] (https://go.dev/doc/install) on how to install it on your computer.

```
go get github.com/vgichira/mpesa-wrapper
```

## Initialize instance of the gateway
To start using the SDK, we first need to create a new instance of the service with the daraja app credentials and the environment. Environment is either SANDBOX or PRODUCTION.

```go
configs, err := mpesa.Init("CONSUMER_KEY_HERE", "CONSUMER_SECRET_HERE", "ENVIROMENT (SANDBOX / LIVE)")
```