package main

import "fmt"

// any number of ineterger arguments
func sum(nums ...int) int {
	tot := 0

	for _, num := range nums {
		tot += num
	}

	return tot
}

// any number of args of any type
func printArgs(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	fmt.Println(1, 2, 3, "Hello") // so on any number of args
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

	slice := []int{1, 2, 3, 4, 5}
	result2 := sum(slice...)
	fmt.Println(result2)

	printArgs(1, 2, "Aman", 'v')
}
