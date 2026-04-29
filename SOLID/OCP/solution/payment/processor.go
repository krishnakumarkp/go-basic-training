package payment

// PaymentMethod defines the contract (abstraction)
// Any new payment method just needs to implement this
type PaymentMethod interface {
	Pay(amount float64)
}

// PaymentProcessor depends on abstraction, not concrete types
type PaymentProcessor struct{}

// ✅ OCP COMPLIANT:
// This function NEVER changes
func (p PaymentProcessor) Process(pm PaymentMethod, amount float64) {
	pm.Pay(amount)
}
