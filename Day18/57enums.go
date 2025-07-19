package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
)

type OrderStatus string

const (
	Pending        OrderStatus = "Pending"
	Shipped                    = "Shipped"
	OutForDelivery             = "Out for Delivery"
	Delivered                  = "Delivered"
	Cancelled                  = "Cancelled"
	Returned                   = "Returned"
	Refunded                   = "Refunded"
	Processing                 = "Processing"
)

func main() {
	var today Day = Friday
	fmt.Println(today)

	var status OrderStatus= Shipped
	fmt.Println(status)
}
