package internal

import "time"

type PaymentResponse struct {
    TransactionID string
    Status        string  
    Amount        float64
    Currency      string
    Timestamp     time.Time
    Provider      string 
}

type PaymentError struct {
    Code    string
    Message string
    Provider string
}