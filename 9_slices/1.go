package main
import "fmt"

func main(){
	// //creating 
	// // var nums[]int --> method 1
	// // nums := []int{} --> method 2
	// //nums := make([]int, 0, 5) -> method 3 --> 3 args, 1st-> type of slice, 2nd -> startign size, 3rd -> capacity
	// nums := make([]int, 0, 5)
	// fmt.Println(nums)
	// fmt.Println("length: ", len(nums))
	// fmt.Println("capacity", cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println("length: ", len(nums))
	// fmt.Println("capacity", cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums);
	// fmt.Println("length: ", len(nums))
	// fmt.Println("capacity", cap(nums))

	//coying slice
	//copy function

	var nums1[]int
	nums1 = append(nums1, 69)
	nums1 = append(nums1, 78)

	nums2 := make([]int, 0, 5)
	nums2 = append(nums2, 99)
	fmt.Println(nums1)
	fmt.Println(nums2)
	copy(nums2, nums1)
	fmt.Println(nums1)
	fmt.Println(nums2)
	
}