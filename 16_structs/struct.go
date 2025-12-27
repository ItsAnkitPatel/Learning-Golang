package main

import (
	"fmt"
	"time"
)

// struct is similar like class concept of oops
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// hacky way of creating constructor in golang
func newOrder(id string, amount float32, status string) *order {
	// any initial setup it goes inside here
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type, (o *order) is the way of connecting function to the struct
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// myOrder := order{
	// 	id:     id,
	// 	amount: amount,
	// 	status: status,
	// }

	myOrder := newOrder("1", 50.00, "received")
	myOrder2 := newOrder("2", 100, "confirmed")
	fmt.Println("order second", myOrder2)

	myOrder.createdAt = time.Now()

	myOrder.changeStatus("confirmed")
	fmt.Println("order struct", myOrder)
	fmt.Println("order amount", myOrder.getAmount())

	// if want to create struct which isn't going to used more than once
	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println("language struct", language)
}
