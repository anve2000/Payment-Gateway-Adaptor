package adapter

import (
    "fmt"
    "payment-gateway-adapter/internal"
    "payment-gateway-adapter/provider"
)

type PaymentAdapter struct {
    provider provider.Provider
}

func NewPaymentAdapter(providerType string) (*PaymentAdapter, error) {
    switch providerType {
    case "A":
        return &PaymentAdapter{provider: provider.NewProviderA()}, nil
    case "B":
        return &PaymentAdapter{provider: provider.NewProviderB()}, nil
    default:
        return nil, fmt.Errorf("unsupported provider: %s", providerType)
    }
}

func (a *PaymentAdapter) ProcessPayment(amount float64, currency string) (*internal.PaymentResponse, error) {
    return a.provider.ProcessPayment(amount, currency)
}