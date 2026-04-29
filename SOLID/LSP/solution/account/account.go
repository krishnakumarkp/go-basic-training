package account

type Account interface {
	GetBalance() float64
}

type Depositable interface {
	Deposit(amount float64)
}

type Withdrawable interface {
	Withdraw(amount float64)
}
