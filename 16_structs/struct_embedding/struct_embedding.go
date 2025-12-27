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
	newOrder.customer.name = "robin"
	fmt.Println(newOrder)
}
