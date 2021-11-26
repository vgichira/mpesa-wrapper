package main

import (
	"testing"

	"github.com/vgichira/mpesa-wrapper"
)

func TestInit(t *testing.T) {
	config := mpesa.Init("pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n", "VCORc4rrhPGP3SRj", "SANDBOX")

	expects := &mpesa.Config{
		ConsumerKey:    "pyhfLWi17bMPs3gchEnEAY2wb6S9Wj9n",
		ConsumerSecret: "VCORc4rrhPGP3SRj",
		Environment:    "SANDBOX",
	}

	if config == nil {
		t.Errorf("got %+v, expected %+v", config, expects)
	}
}
