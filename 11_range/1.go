package main

import "fmt"

// iterating over structures
func main() {
	// nums := []int{6, 7, 8}

	// sum := 0

	// //index, value := range nums
	// for _, num := range nums{
	// 	sum += num
	// }

	// fmt.Println(sum)

	// //in case of map

	// m := map[string]string{"fname":"John", "sname":"Doe"}

	// for k, v := range m{
	// 	fmt.Println(k, v)
	// }

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

	//i is the index and c is the unicode code of each character
	//unicode of each character can go beyond 255 as well (unlike ASCII)
	//hence we might end up getting 2 bytes per character instead of 1 byte like we have in ASCII

}
