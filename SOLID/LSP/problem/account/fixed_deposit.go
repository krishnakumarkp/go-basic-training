package account

import "fmt"

type FixedDepositAccount struct {
	Balance float64
}

func (f *FixedDepositAccount) Deposit(amount float64) {
	f.Balance += amount
	fmt.Println("Deposited to FD:", amount)
}

// ❌ Forced to implement Withdraw → violates LSP
func (f *FixedDepositAccount) Withdraw(amount float64) {
	panic("withdrawal not allowed before maturity")
}
