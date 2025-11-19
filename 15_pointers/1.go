package main

import "fmt"

// be reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("value in chnageNum: ", *num)
}
func main() {
	num := 10
	changeNum(&num)
	println("Value in main: ", num)

}
