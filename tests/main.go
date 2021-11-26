package main

import (
	"fmt"

	"github.com/vgichira/mpesa-wrapper"
)

func main() {
	service := mpesa.Init("pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n", "VCORc4rrhPGP3SRj", "SANDBOX")

	// STK Push Request Body
	requestBody := &mpesa.C2BTransaction{
		CommandID:     "CustomerPayBillOnline",
		Amount:        "100",
		MSISDN:        "254725089232",
		BillRefNumber: "",
		ShortCode:     "",
	}

	resp, err := service.SimulateC2B(requestBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
