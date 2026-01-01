package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  // struct embedding
}

// NOTE: Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
func main() {
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
		// customer:newCustomer,
	}
	// newOrder.customer.name = "robin"
	newOrder.name = "robin"
	// fmt.Println(newOrder)
	fmt.Println(newOrder.name)
}
