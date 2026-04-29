package paypal

import "fmt"

type PayPal struct{}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("Processing PayPal payment of %.2f\n", amount)
}
