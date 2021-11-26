package main

import (
	"fmt"

	"github.com/vgichira/mpesa-wrapper"
)

func main() {
	service := mpesa.Init("pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n", "VCORc4rrhPGP3SRj", "SANDBOX")

	// STK Push Request Body
	requestBody := &mpesa.Reversal{
		Initiator:              "TestG2Init",
		SecurityCredential:     "EsJocK7+NjqZPC3I3EO+TbvS+xVb9TymWwaKABoaZr/Z/n0UysSsEfea4eQyeWWmyx0t7K1vmfUlGk",
		CommandID:              "BusinessPayment",
		TransactionID:          "",
		Amount:                 "10",
		ReceiverParty:          "",
		ReceiverIdentifierType: "",
		Remarks:                "here are my remarks",
		QueueTimeoutURL:        "https://mydomain.com/path",
		ResultURL:              "https://mydomain.com/path",
		Occasion:               "Christmas",
	}

	resp, err := service.ReverseTransaction(requestBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
