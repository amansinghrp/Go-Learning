// package main

// import "fmt"

// // payment struct
// type payment struct {
// 	gateway razorpay
// }

// // violates open close principle -> methods should be open for extension but closed for modification
// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGw := razorpay{}
// 	// razorpayPaymentGw.pay(amount)

// 	// stripePaymentGw := stripe{}
// 	// stripePaymentGw.pay(amount)

// 	p.gateway.pay(amount)
// }

// // razorpay payment gateway
// type razorpay struct {
// }

// func (r razorpay) pay(amount float32) {
// 	//logic to make payment
// 	fmt.Println("Making payment using razorpay", amount)
// }

// // stripe payment gateway
// type stripe struct {
// }

// func (s stripe) pay(amount float32) {
// 	fmt.Println("Made payment using stripe: ", amount)
// }

// type fakePayment struct {
// }

// func (f fakePayment) pay(amount float32) {
// 	fmt.Println("Makiing payment using fake gateway: ", amount)

// }

// func main() {
// 	// myPayment := payment{}
// 	// stripePaymentGw := stripe{}
// 	razorpayPaymentGw := razorpay{}
// 	// fakePaymentGw := fakePayment{}
// 	myPayment := payment{
// 		// gateway: stripePaymentGw,
// 		gateway: razorpayPaymentGw,
// 		// gateway: fakePaymentGw, ---> cannot do this coz payment struct has razorpay payment gateway as defined

// 	}
// 	myPayment.makePayment(100)
// }