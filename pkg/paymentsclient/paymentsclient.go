package paymentsclient

import (
	"context"
	"errors"
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
func (a PaymentsClient) MakePayment(amount int64, from string, to string) (pb.Payment, error) {
	// make the payment request
	paymentRequest := pb.PaymentRequest{
		Amount: amount,
		From:   from,
		To:     to,
	}

	payment, err := a.externalClient.Pay(context.Background(), &paymentRequest)

	if err != nil {
		log.Println(err)

		return pb.Payment{}, errors.New("Failed to send payment")
	}

	return *payment, nil
}

// ConfirmPayment confirms a payment
func (a PaymentsClient) ConfirmPayment(id string) (pb.Payment, error) {
	paymentByID := pb.PaymentByIdRequest{
		Id: id,
	}

	confirm, err := a.externalClient.Confirm(context.Background(), &paymentByID)

	if err != nil {
		return pb.Payment{}, errors.New("Failed to confirm payment")
	}

	return *confirm, nil
}

// GetByID gets a specific payment by id
func (a PaymentsClient) GetByID(id string) (pb.Payment, error) {
	paymentByID := pb.PaymentByIdRequest{
		Id: id,
	}
	payment, err := a.externalClient.GetById(context.Background(), &paymentByID)

	if err != nil {
		log.Panicln(err)
		return pb.Payment{}, errors.New("Failed to get payment by id")
	}
	return *payment, nil
}

// GetAll gets all payments made through service
func (a PaymentsClient) GetAll() (pb.PaymentList, error) {
	all := pb.AllPaymentsRequest{}

	payments, err := a.externalClient.GetAll(context.Background(), &all)

	if err != nil {
		log.Println(err)
		return pb.PaymentList{}, errors.New("Failed to get all payments")
	}

	return *payments, nil
}
