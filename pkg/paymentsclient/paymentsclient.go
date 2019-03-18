package paymentsclient

import (
	"log"

	pb "github.com/kimpettersen/svc-payments/proto"
	"google.golang.org/grpc"
)

// PaymentsClient is responsible for calling svc-payments
type PaymentsClient struct {
	externalClient pb.PaymentsClient
}

// New is a constructor for a PaymentsClient. We don't have a special method for this in Go
func New() *PaymentsClient {
	address := "127.0.0.1:3000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %+v", err)
	}
	client := pb.NewPaymentsClient(conn)

	return &PaymentsClient{
		externalClient: client,
	}
}

// MakePayment calls the external gRPC endpoint to make a payment
func (a PaymentsClient) MakePayment() {
	// TODO: Make a remote call to the payment server and make a payment
}
