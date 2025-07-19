package main

import "fmt"

type Customer struct {
	name string
	age  int
	city string
}

type Orders struct {
	orderId   int
	orderName string
	orderCost int
	Customer  // Embedding Customer struct
}

func main() {
	order1 := Orders{
		orderId:   1,
		orderName: "Laptop",
		orderCost: 5000,
		Customer: Customer{
			name: "Amit Khurana",
			age:  18,
			city: "Delhi",
		},
	}
	fmt.Println(order1.orderCost)
}