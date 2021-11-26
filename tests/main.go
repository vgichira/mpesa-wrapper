package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/vgichira/mpesa-wrapper"
)

func main() {
	service := mpesa.Init("pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n", "VCORc4rrhPGP3SRj", "LIVE")

	timestamp := time.Now().Format("20060102150405")

	// STK Push Request Body
	requestBody := &mpesa.LipaNaMpesaRequest{
		BusinessShortCode: "174379",
		Password:          base64.StdEncoding.EncodeToString([]byte("174379" + "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919" + timestamp)),
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            "254725089232",
		PartyB:            "174379",
		PhoneNumber:       "254725089232",
		CallBackURL:       "https://google.com",
		AccountReference:  "Test",
		TransactionDesc:   "Test",
	}

	resp, err := service.LipaNaMpesaOnline(requestBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
