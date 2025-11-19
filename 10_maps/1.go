package main

import (
	"fmt"
	"maps"
)

func main() {
	// //creating maps:
	// //m := make(map[key type]value type)
	// m := make(map[string]string)

	// //setting elements
	// m["name"] = "Go lang"
	// m["author"] = "Aman"

	// //get elements
	// fmt.Println(m["name"], m["author"])

	// //IMP: if I try to access a key which doesnot exist in the map
	// 	//then it returns zeroed value

	// m := make(map[string]int)

	// m["age"] = 40
	// m["price"] = 60
	// fmt.Println(m["aman"])//prints zeroed value of int


	// fmt.Println(len(m))

	// fmt.Println(m)

	// //to delete and element
	// delete(m, "age")

	// fmt.Println(m);


	// //to clear the complete map
	// clear(m)

	// fmt.Println(m)

	// //check if element is present or not:
	// val, ok := m["hgh"]

	// if ok {
	// 	fmt.Println("Element is present", val)
	// }else{
	// 	fmt.Println("Element is not present", val)
	// }



	m1 := map[string]int{"age" : 40, "size": 34}
	m2 := map[string]int{"age":40, "size" :34}

	fmt.Println(maps.Equal(m1, m2))
}
