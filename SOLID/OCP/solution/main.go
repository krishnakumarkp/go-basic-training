package main

import (
	"krishnakumarkp/ocp/solution/creditcard"
	"krishnakumarkp/ocp/solution/payment"
	"krishnakumarkp/ocp/solution/paypal"
	"krishnakumarkp/ocp/solution/upi"
)

func main() {
	processor := payment.PaymentProcessor{}

	cc := creditcard.CreditCard{}
	pp := paypal.PayPal{}

	processor.Process(cc, 100)
	processor.Process(pp, 200)

	u := upi.UPI{}
	processor.Process(u, 300)
}
