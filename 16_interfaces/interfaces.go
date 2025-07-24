package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float64) {
	razorpayGw := razorpay{}
	razorpayGw.pay(amount)
	fmt.Println("Payment processed:", amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64) {
	fmt.Println("Payment made using Razorpay:", amount)
}

func main() {

	p := payment{}
	p.makePayment(100.50)
}