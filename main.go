package main

import (
	"github.com/kimpettersen/svc-voicepayments/pkg/paymentsclient"
)

func main() {
	paymentsClient := paymentsclient.New()
	paymentsClient.MakePayment()
}
