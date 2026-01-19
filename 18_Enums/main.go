package main

import "fmt"

type OrderStatus int

const (
	Recived OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

// enumerated types
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
