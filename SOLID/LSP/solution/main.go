package main

import "krishnakumarkp/lsp/solution/account"

func ProcessDeposit(d account.Depositable) {
	d.Deposit(500)
}

func ProcessWithdrawal(w account.Withdrawable) {
	w.Withdraw(200)
}

func main() {
	sa := &account.SavingsAccount{Balance: 1000}
	fd := &account.FixedDepositAccount{Balance: 1000}

	// ✅ Both support deposit
	ProcessDeposit(sa)
	ProcessDeposit(fd)

	// ✅ Only savings supports withdrawal
	ProcessWithdrawal(sa)

	// ProcessWithdrawal(fd)
	// ❌ Compile-time error → correct behavior
}
