// internal/model.go
package internal

import "fmt"
import "time"

// PaymentResponse is the normalized internal response structure
type PaymentResponse struct {
    TransactionID string    `json:"transaction_id"`
    Status        string    `json:"status"`     // APPROVED, FAILED, PENDING
    Amount        float64   `json:"amount"`
    Currency      string    `json:"currency"`
    Timestamp     time.Time `json:"timestamp"`
    Provider      string    `json:"provider"`   // To track which provider was used
}

type PaymentError struct {
    Code     string `json:"code"`
    Message  string `json:"message"`
    Provider string `json:"provider"`
}

func (e *PaymentError) Error() string {
    return fmt.Sprintf("[%s] %s (Provider: %s)", e.Code, e.Message, e.Provider)
}