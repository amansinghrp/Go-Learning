package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	//every struct that has this pay method automatically
	//implements this method. and uses the impleentation of the
	//methods provided in it.
	//implicitly decided by go, just the function signature shoudl be same
}

// payment struct
type payment struct {
	gateway paymenter
}

// violates open close principle -> methods should be open for extension but closed for modification
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

// razorpay payment gateway
type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

// stripe payment gateway
type stripe struct {
}

func (s stripe) pay(amount float32) {
	fmt.Println("Made payment using stripe: ", amount)
}

type fakePayment struct {
}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Makiing payment using fake gateway: ", amount)

}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	fakePaymentGw := fakePayment{}
	myPayment := payment{
		// gateway: stripePaymentGw,
		// gateway: razorpayPaymentGw,
		gateway: fakePaymentGw, //---> now we can easily do this because of interface

	}
	myPayment.makePayment(100)
}
