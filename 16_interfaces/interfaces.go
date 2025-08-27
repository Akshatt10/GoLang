package main

import "fmt"

type payment struct {
	// We can add more fields here if needed
	gateway paymenter
}

type paymenter interface {
	pay(amount float64)
}

func (p payment) makePayment(amount float64) {
	// razorpayGw := razorpay{}
	// razorpayGw.pay(amount)

	// stripeGw := stripe{}
	p.gateway.pay(amount)

	// We are voilating the open close principle here
	// because if we want to add a new payment gateway,
	// we need to modify the makePayment method.
	// This is not a good design.

	fmt.Println("Payment processed:", amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64) {
	fmt.Println("Payment made using Razorpay:", amount)
}

type stripe struct{}

type paypal struct{}

func (p paypal) pay(amount float64) {
	fmt.Println("Makeing payment through paypal:", amount)
}

func (s stripe) pay(amount float64) {
	fmt.Println("Makeing payment through stripe:", amount)
}

func main() {
	paypalGw := paypal{}
	
	newpayment := payment{
		gateway: paypalGw,
	}
	newpayment.makePayment(2000)
}
