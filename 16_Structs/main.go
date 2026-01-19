package main

import (
	"fmt"
	"time"
)

// in Go, structs are used to group related data together like classes in other languages

type customer struct {
	name    string
	age     int
	address string
}

// struct with embedded struct
type Order struct {
	ID        string
	Amount    float32
	Status    string
	createdAt time.Time
	customer
}

// func newOrder(id string, amount float32, status string) *Order {

// 	// initial setup
// 	Order := Order{
// 		ID:     id,
// 		Amount: amount,
// 		Status: status,
// 	}
// 	return &Order
// }

// reciver type
func (o *Order) changeStatus(status string) {
	o.Status = status
}

func main() {

	order := Order{
		ID:        "12345",
		Amount:    250.75,
		Status:    "Pending",
		createdAt: time.Now(),
		customer: customer{
			name:    "krishna",
			age:     21,
			address: "123 Main St",
		},
	}

	order.changeStatus("Shipped")

	fmt.Println("Order Details:")
	fmt.Println("ID:", order.ID)
	fmt.Println("Amount:", order.Amount)
	fmt.Println("Status:", order.Status)
	fmt.Println("Created At:", order.createdAt)
	fmt.Println("Customer Name:", order.customer.name)
	fmt.Println("Customer Age:", order.customer.age)
	fmt.Println("Customer Address:", order.customer.address)
	fmt.Println(order)

	language := struct {
		name   string
		isGood bool
	}{
		name:   "krishna",
		isGood: true,
	}
	fmt.Println(language)
}
