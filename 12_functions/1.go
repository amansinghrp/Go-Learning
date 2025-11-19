package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "C++", "Python", "Golang"
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4 * a
	}

}

func main() {
	// fmt.Println(add(5, 7))

	// lang1, lang2, _ := getLanguages()
	// fmt.Println(getLanguages())
	// fmt.Println(lang1, lang2)

	//passing functions inside function
	fn := processIt()
	fmt.Println(fn(4))

	fmt.Println(processIt()(4))
}
