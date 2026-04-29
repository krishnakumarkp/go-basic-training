package creditcard

import "fmt"

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Processing credit card payment of %.2f\n", amount)
}
