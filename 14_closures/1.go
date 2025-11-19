package main

import "fmt"

func counter() func() int {
	cnt := 0

	return func() int {
		cnt++
		return cnt //this variable will be stored in the closure environment -> heap
	}
}

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	//eg1
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	//eg 2
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))

}
