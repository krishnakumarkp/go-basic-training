package main

import "krishnakumarkp/lsp/problem/account"

func ProcessAccount(a account.Account) {
	a.Deposit(500)
	a.Withdraw(200) // assumes all accounts can withdraw
}

func main() {
	sa := &account.SavingsAccount{Balance: 1000}
	fd := &account.FixedDepositAccount{Balance: 1000}

	ProcessAccount(sa) // ✅ works
	ProcessAccount(fd) // 💥 runtime panic
}
