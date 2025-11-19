package main

import (
	"fmt"
	"time"
)

// for struct embeding
type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
	customer
}

// workaround for making constructors in struct in go lang
// kind of like -> ek function hi ho gya alag se jo bas uss type ka object bana ke
// vapas uska reference return kr deta hai
// -> cool
func newOrder(id string, amount float64, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

//receiver type

// func(o order) -> is accoridng to convention to associate methods to struct
func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

func (o order) getAmount() float64 {
	return float64(o.amount)
}

func main() {
	// orders1 := order{
	// 	id:     "101",
	// 	amount: 101.90,
	// 	status: "Ordered",
	// }

	// orders1 := newOrder("101", 105.89, "ordered")

	// orders1.createdAt = time.Now()
	// fmt.Println(orders1)
	// orders1.changeStatus("Shipped")
	// fmt.Println("After chanigng status:")
	// fmt.Println(orders1)

	// fmt.Println("The amount was: ", orders1.getAmount())

	// //direct initilaisation
	// langauge := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(langauge)

	newCustomer := customer{
		name:  "Aman",
		phone: "90876541231",
	}
	orders1 := order{
		id:     "101",
		amount: 101.90,
		status: "Ordered",
		// customer: customer{
		// 	name:  "Aman",
		// 	phone: "012345567789",
		// },
		customer: newCustomer,
	}

	fmt.Println(orders1)
}
