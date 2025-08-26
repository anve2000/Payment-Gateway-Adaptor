package provider

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"

    "payment-gateway-adapter/internal"
)

type ProviderA struct{}

func NewProviderA() Provider {
    return &ProviderA{}
}

func (p *ProviderA) ProcessPayment(amount float64, currency string) (*internal.PaymentResponse, *internal.PaymentError) {
    if rand.Float64() < 0.8 {
        resp := map[string]interface{}{
            "transaction_id": fmt.Sprintf("TXN%d", time.Now().UnixNano()),
            "status":         "APPROVED",
            "amount":         int(amount * 100), 
            "currency":       currency,
            "timestamp":      time.Now().UTC().Format(time.RFC3339),
        }
        jsonData, _ := json.Marshal(resp)

        var result struct {
            TransactionID string  `json:"transaction_id"`
            Status        string  `json:"status"`
            Amount        int     `json:"amount"`
            Currency      string  `json:"currency"`
            Timestamp     string  `json:"timestamp"`
        }
        json.Unmarshal(jsonData, &result)

        ts, _ := time.Parse(time.RFC3339, result.Timestamp)

        return &internal.PaymentResponse{
            TransactionID: result.TransactionID,
            Status:        result.Status,
            Amount:        float64(result.Amount) / 100,
            Currency:      result.Currency,
            Timestamp:     ts,
            Provider:      "ProviderA",
        }, nil
    } else {
        return nil, &internal.PaymentError{
            Code:     "INSUFFICIENT_FUNDS",
            Message:  "Not enough balance",
            Provider: "ProviderA",
        }
    }
}