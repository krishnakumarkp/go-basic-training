package upi

import "fmt"

type UPI struct{}

func (u UPI) Pay(amount float64) {
	fmt.Printf("Processing UPI payment of %.2f\n", amount)
}
