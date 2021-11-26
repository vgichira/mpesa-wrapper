package main

import (
	"fmt"

	"github.com/vgichira/mpesa-wrapper"
)

func main() {
	service := mpesa.Init("pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n", "VCORc4rrhPGP3SRj", "SANDBOX")

	// STK Push Request Body
	requestBody := &mpesa.RegisterURL{
		ValidationURL:   "https://google.com/",
		ConfirmationURL: "https://google.com/",
		ResponseType:    "Completed",
		ShortCode:       "174379",
	}

	resp, err := service.RegisterURL(requestBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
