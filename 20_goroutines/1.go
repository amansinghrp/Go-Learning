package main

import (
	"fmt"
	"sync"
)

// func task(id int) {
// 	fmt.Println("Doing task: ", id)
// }

// prblem -> we are deliberately pausing main function so that others can fiinish execution
// func main() {
// 	for i := 0; i < 10; i++ {
// 		go task(i)

// 		// //inline functions with go routines
// 		// go func(i int){
// 		// 	fmt.Println(i)
// 		// }(i)
// 	}

// 	time.Sleep(time.Second * 1) //but in real world use cases we cant use this
// 	//hence we use waitgroups
// }

//actually we shoudl be using waitgroups

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer keyword specifies that execute this only. after function has finished execution
	fmt.Println("Doing task: ", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // add one waitgroup
		go task(i, &wg)
	}

	wg.Wait()

}
