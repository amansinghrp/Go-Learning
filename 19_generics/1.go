package main

import "fmt"

// func printSlices[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Print(item, " ")
// 	}
// 	fmt.Println()
// }

// func printSlices[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Print(item, " ")
// 	}
// 	fmt.Println()
// }

func printSlices[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Print(item, " ", name)
	}
	fmt.Println()
}

func main() {
	nums := []int{1, 2, 3, 4}
	nums2 := []string{"Aman", "Amar", "Keerat"}
	// nums3 := []float32{3.9, 8, 9}
	printSlices(nums, "nalla")
	printSlices(nums2, "basanti")
	// printSlices(nums3)

}
