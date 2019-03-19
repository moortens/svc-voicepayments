package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kimpettersen/svc-voicepayments/pkg/paymentsclient"
)

func main() {
	get := flag.String("get", "", "get specific payment by id")
	confirm := flag.String("confirm", "", "id of payment to confirm")
	all := flag.Bool("all", false, "get all payments")

	to := flag.String("to", "", "user to send payment to")
	from := flag.String("from", "", "user to send payment from")
	amount := flag.Int64("amount", 0, "amount to send")

	flag.Parse()

	paymentsClient := paymentsclient.New()

	if *get != "" {
		payment, err := paymentsClient.GetByID(*get)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Payment from %s to %s of amount %d\n",
			payment.From,
			payment.To,
			payment.Amount)

		return
	}

	if *confirm != "" {
		payment, err := paymentsClient.ConfirmPayment(*confirm)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Payment of id %s confirmed", payment.Id)

		return
	}

	if *all == true {
		payments, err := paymentsClient.GetAll()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		for i := 0; i < len(payments.Payments); i++ {
			fmt.Printf("%d. Payment from %s to %s of amount %d (%s)\n",
				i,
				payments.Payments[i].From,
				payments.Payments[i].To,
				payments.Payments[i].Amount,
				payments.Payments[i].Id)
		}

		return
	}

	if *to != "" && *from != "" && *amount != 0 {
		payment, err := paymentsClient.MakePayment(*amount, *from, *to)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Please confirm payment: %s\n", payment.Id)

		return
	}

	flag.Usage()
}
