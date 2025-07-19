package main

import "fmt"

type paymentInterface interface {
	pay(amount int)
}

// open-close pronciple
// we can extend the functionality of the code without modifying the existing code
// we can add new payment gateways without changing the existing code
type gateway struct {
	gatewayName paymentInterface
}

type razorpay struct{}

func (r razorpay) pay(amount int) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (r stripe) pay(amount int) {
	fmt.Println("making payment using stripe", amount)
}

type paypal struct{}

func (r paypal) pay(amount int) {
	fmt.Println("making payment using paypal", amount)
}
func main() {
	gateway := gateway{
		gatewayName: paypal{},
	}

	gateway.gatewayName.pay(1000)
}
