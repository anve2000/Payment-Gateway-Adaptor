// main.go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"payment-gateway-adapter/adapter"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Test both providers
	for _, provider := range []string{"A", "B"} {
		fmt.Printf("\n--- Testing %s ---\n", provider)
		adapter, err := adapter.NewPaymentAdapter(provider)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		amount := 99.99
		currency := "USD"

		resp, err := adapter.ProcessPayment(amount, currency)
		if err != nil {
			fmt.Printf("Payment failed: %v\n", err) // Uses .Error() method
		} else {
			fmt.Printf("Payment successful!\n")
			fmt.Printf("ID: %s | Amount: %.2f %s | Time: %v | Provider: %s\n",
				resp.TransactionID, resp.Amount, resp.Currency, resp.Timestamp, resp.Provider)
		}
	}
}
