package main

import "krishnakumarkp/ocp/problem/payment"

func main() {
	processor := payment.PaymentProcessor{}

	processor.Process("credit", 100)
	processor.Process("paypal", 200)
}
