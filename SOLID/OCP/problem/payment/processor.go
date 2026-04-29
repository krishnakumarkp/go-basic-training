package payment

import "fmt"

// PaymentProcessor handles multiple payment types
type PaymentProcessor struct{}

// Violates OCP:
// Every time a new payment method is introduced,
// this function must be modified
func (p PaymentProcessor) Process(paymentType string, amount float64) {
	if paymentType == "credit" {
		fmt.Printf("Processing credit card payment of %.2f\n", amount)
	} else if paymentType == "paypal" {
		fmt.Printf("Processing PayPal payment of %.2f\n", amount)
	} else if paymentType == "debit" {
		fmt.Printf("Processing debit card payment of %.2f\n", amount)
	} else {
		fmt.Println("Unsupported payment method")
	}
}
