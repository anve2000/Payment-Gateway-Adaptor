package provider

import "payment-gateway-adapter/internal"

type Provider interface {
    ProcessPayment(amount float64, currency string) (*internal.PaymentResponse, *internal.PaymentError)
}