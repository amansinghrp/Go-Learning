package main

import "fmt"

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeStatus(nstatus OrderStatus) {
	fmt.Println("Changing order status to: ", nstatus)
}

func main() {
	changeStatus(Received)
	changeStatus(Confirmed)
	changeStatus(Prepared)
}
