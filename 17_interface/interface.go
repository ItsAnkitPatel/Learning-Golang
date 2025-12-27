package main

import (
	"fmt"
)

// NOTE: convention for writing interface: usually we add 'r' at the end of each interface name
// store => storer, payment=> paymentr
// Do not compare typescript interface with golang interface
// IMPORTANT: Read about 'dependency inversion'
type paymentr interface {
	pay(amount float32) //method
}

type payment struct {
	gateway paymentr
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

func main() {
	razorpayGw := razorpay{}
	newPayment := payment{
		gateway: razorpayGw,
	}
	newPayment.makePayment(100)
}
