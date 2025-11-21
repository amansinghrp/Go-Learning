package main

import (
	"fmt"
	"time"
)

// func processNum(ch chan int) {
// 	fmt.Println("Processing number: ", <-ch)
// }

// func main() {

// 	numChannel := make(chan int)

// 	go processNum(numChannel) // starts another go routine -> receiver

// 	numChannel <- 10 // the current main function (another goroutine) simply sends data

// 	time.Sleep(time.Second)
// 	// tanishk := make(chan int)

// 	// tanishk <- 5 //blocking nature, doesnot sends data untill someone is receiving

// 	// msg := <-tanishk // blocking nature -> doesnot receives data untill soemone is sending
// 	// //hence deadlock

// 	// fmt.Println(msg)

// }

// //sending data in the form a queue
// func processNum(ch chan int) {
// 	//instead of receiving a single data, we loop over a queue in the channel
// 	// fmt.Println("Processing number: ", <-ch)

// 	for num := range ch {
// 		fmt.Println("Processing number: ", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {

// 	numChannel := make(chan int)

// 	go processNum(numChannel) // starts another go routine -> receiver

// 	// numChannel <- 10 // the current main function (another goroutine) simply sends data
// 	for {
// 		numChannel <- rand.IntN(100)
// 	}

// 	// time.Sleep(time.Second)

// }

// // receiving data in the form a queue
// func sum(total chan int, num1 int, num2 int) {
// 	ans := num1 + num2
// 	total <- ans
// }

// func main() {
// 	sumChannel := make(chan int)

// 	go sum(sumChannel, 3, 6)

// 	ans := <-sumChannel // blockig, therefore we did not use sleep in the main functions

// 	fmt.Println(ans)

// }

// // impleting the waitgroup like logic for synchronization using channels
// func task(done chan bool) {
// 	defer func() {
// 		done <- true //even sending false wont make a difference as we are more concerned about sending some data rather than what that data is
// 	}()
// 	fmt.Println("Processing...")
// }

// func main() {
// 	doneChannel := make(chan bool)

// 	go task(doneChannel)

// 	<-doneChannel // add a blocking receiving end so that the main function does not terminate early
// 	//it will unblock once it receives any value through the channel, be it be true or false

// }

// //buffered channels

// func main() {
// 	emailChannel := make(chan string, 2) //buffer size of 2

// 	emailChannel <- "aman@mmt.com"
// 	emailChannel <- "someone@mmt.com"
// 	// emailChannel <- "test@mmt.com" // this will block as the buffer size is only 2

// 	fmt.Println(<-emailChannel)
// 	fmt.Println(<-emailChannel)
// 	// fmt.Println(<-emailChannel)// this will block the process and will be in deadlock coz now the buffer is empty after receiving 2 values

// }

// // implementing quueue system for sending data
// func emailSender(emailChannel chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	//block operation
// 	for email := range emailChannel {
// 		fmt.Println("Sending email to: ", email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	emailChannel := make(chan string, 10)
// 	done := make(chan bool)
// 	go emailSender(emailChannel, done)
// 	for i := 0; i < 5; i++ {
// 		emailChannel <- fmt.Sprintf("employee_%d@gmail.com", i)
// 	}
// 	fmt.Println("Done sending emails")

// 	//this is important as the range over channel is a blocking operation
// 	close(emailChannel)
// 	<-done
// }

// // listening upon multiple channels --> select keyword use
// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan string)

// 	go func() {
// 		chan2 <- "pong"
// 	}()
// 	go func() {
// 		chan1 <- 10
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case chan1Val := <-chan1:
// 			fmt.Println("Received data from channel 1: ", chan1Val)
// 		case chan2Val := <-chan2:
// 			fmt.Println("Received data from channel 2: ", chan2Val)
// 		}
// 	}
// }

// stricly ensuring that a particular channel is for send only or receice only
// ensures type safety
func emailSender(emailChannel <-chan string, done chan<- bool) {
	// emailChannel <-chan string. --> arrow outwards channel : send only channel
	//done chan<- bool --> arrow towrads channel : receive only channel
	defer func() { done <- true }()

	//block operation
	for email := range emailChannel {
		fmt.Println("Sending email to: ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChannel := make(chan string, 10)
	done := make(chan bool)
	go emailSender(emailChannel, done)
	for i := 0; i < 5; i++ {
		emailChannel <- fmt.Sprintf("employee_%d@gmail.com", i)
	}
	fmt.Println("Done sending emails")

	//this is important as the range over channel is a blocking operation
	close(emailChannel)
	<-done
}
