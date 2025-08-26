# Payment Gateway Adapter System

A robust, extensible Go-based payment gateway adapter that normalizes interactions with multiple payment providers into a unified interface.

Built to handle varying response formats, error structures, and operational behaviors - while being easily extendable for future providers.


## Features

- Unified API: Single interface to process payments across providers
- Response Normalization: Diverse provider responses mapped to a consistent internal model
- Error Standardization: Provider-specific errors converted to a common `Error` interface with structured metadata
- Extensible Design: Add new providers via interface implementation
- Production-Ready Patterns: Adapter pattern, interface abstraction, proper error handling


## How to Add a New Provider
To add a new payment provider, follow these steps:

1. Create a new file in the `provider` directory (e.g., `provider_c.go`).
2. Define a struct for the provider (e.g., `ProviderC`).
3. Implement the `ProcessPayment(amount float64, currency string) (*internal.PaymentResponse, error)` method. This method should:
   - Simulate or call the external provider's API
   - Parse the response
   - Map the response to the `internal.PaymentResponse` struct on success
   - Return a formatted error using `internal.PaymentError` on failure
4. Register the new provider in `adapter/adapter.go` by adding a case for it in the `NewPaymentAdapter` function.
5. Use the provider by passing its name (e.g., "C") when creating the adapter.


```git clone https://github.com/your-username/payment-gateway-adapter.git
cd payment-gateway-adapter