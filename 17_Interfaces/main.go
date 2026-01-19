package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// open close principle violated
func (p payment) makePayment(amount float32) {
	// razorpayGw := razorpay{}
	// razorpayGw.pay(amount)
	// paypalGw := paypal{}
	// paypalGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to pay
	fmt.Println("Paid using Razorpay:", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	// logic to pay
	fmt.Println("Paid using PayPal:", amount)
}

func main() {
	payment1 := paypal{}
	// payment2 := razorpay{}
	newPay := payment{gateway: payment1}
	// newPay := payment{}
	newPay.makePayment(1000)
}
