package provider

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"
    "payment-gateway-adapter/internal"
)

type ProviderB struct{}

func NewProviderB() Provider {
    return &ProviderB{}
}

func (p *ProviderB) ProcessPayment(amount float64, currency string) (*internal.PaymentResponse, *internal.PaymentError) {
    if rand.Float64() < 0.7 {
        timestamp := time.Now().Unix()
        resp := map[string]interface{}{
            "paymentId": fmt.Sprintf("PAY-%d-%s", timestamp, "ABC"),
            "state":     "SUCCESS",
            "value": map[string]string{
                "amount":       fmt.Sprintf("%.2f", amount),
                "currencyCode": currency,
            },
            "processedAt": timestamp,
        }
        jsonData, _ := json.Marshal(resp)

        var result struct {
            PaymentID   string            `json:"paymentId"`
            State       string            `json:"state"`
            Value       map[string]string `json:"value"`
            ProcessedAt int64             `json:"processedAt"`
        }
        json.Unmarshal(jsonData, &result)

        ts := time.Unix(result.ProcessedAt, 0)

        amountVal := amount
        fmt.Sscanf(result.Value["amount"], "%f", &amountVal)

        return &internal.PaymentResponse{
            TransactionID: result.PaymentID,
            Status:        "APPROVED", // normalize to common status
            Amount:        amountVal,
            Currency:      result.Value["currencyCode"],
            Timestamp:     ts,
            Provider:      "ProviderB",
        }, nil
    } else {
        return nil, &internal.PaymentError{
            Code:     "PAYMENT_FAILED",
            Message:  "Card declined",
            Provider: "ProviderB",
        }
    }
}