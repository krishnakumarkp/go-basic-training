package account

import "fmt"

type SavingsAccount struct {
	Balance float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
	fmt.Println("Deposited:", amount)
}

func (s *SavingsAccount) Withdraw(amount float64) {
	s.Balance -= amount
	fmt.Println("Withdrawn:", amount)
}
