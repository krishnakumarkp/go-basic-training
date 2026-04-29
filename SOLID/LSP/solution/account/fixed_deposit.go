package account

import "fmt"

type FixedDepositAccount struct {
	Balance float64
}

func (f *FixedDepositAccount) GetBalance() float64 {
	return f.Balance
}

// ✅ Supports deposit
func (f *FixedDepositAccount) Deposit(amount float64) {
	f.Balance += amount
	fmt.Println("Deposited to FD:", amount)
}

// ❌ No Withdraw → still LSP compliant
