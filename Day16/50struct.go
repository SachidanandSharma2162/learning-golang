package main

import (
	"fmt"
	"time"
)

type Order struct {
	orderId      int
	customerName string
	status       string
	orderAT      time.Time
	cost         int
}

// call by reference function for order status change
func (O *Order) changeOrderStatus() {
	if O.status == "Pending" {
		O.status = "Shipped"
	} else if O.status == "Shipped" {
		O.status = "Delivered"
	}
}
func (O *Order) getOrderPrice(status string) {
	O.status = status
}
func main() {
	order1 := Order{
		orderId:      1,
		customerName: "John Doe",
		status:       "Pending",
		orderAT:      time.Now(),
		cost:         1000,
	}
	order2 := Order{
		orderId:      2,
		customerName: "Snow Doe",
		status:       "Shipped",
		orderAT:      time.Now(),
		cost:         2000,
	}
	fmt.Println(order1.orderId, order1.customerName, order1.status, order1.orderAT, order1.cost)
	order1.changeOrderStatus()
	fmt.Println(order1.orderId, order1.customerName, order1.status, order1.orderAT, order1.cost)
	fmt.Println(order1.cost)
	fmt.Println(order2.orderId, order2.customerName, order2.status, order2.orderAT, order2.cost)
	order2.changeOrderStatus()
	fmt.Println(order2.orderId, order2.customerName, order2.status, order2.orderAT, order2.cost)

}
