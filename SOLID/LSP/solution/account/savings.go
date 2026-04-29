package account

import "fmt"

type SavingsAccount struct {
	Balance float64
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.Balance
}

// ✅ Supports deposit
func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
	fmt.Println("Deposited:", amount)
}

// ✅ Supports withdrawal
func (s *SavingsAccount) Withdraw(amount float64) {
	s.Balance -= amount
	fmt.Println("Withdrawn:", amount)
}
