package handlers

import (
	"context"

	paymentv1 "github.com/ol1mov-dev/protos/pkg/payment/v1"
)

type PaymentServer struct {
	paymentv1.UnsafePaymentServiceServer
}

func (s *PaymentServer) CreatePayment(ctx context.Context, req *paymentv1.CreatePaymentRequest) (*paymentv1.CreatePaymentResponse, error) {
	return &paymentv1.CreatePaymentResponse{}, nil
}

func (s *PaymentServer) GetPayment(ctx context.Context, req *paymentv1.GetPaymentRequest) (*paymentv1.GetPaymentResponse, error) {
	return &paymentv1.GetPaymentResponse{}, nil
}
