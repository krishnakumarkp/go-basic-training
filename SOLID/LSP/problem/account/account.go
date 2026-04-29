package account

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
}
