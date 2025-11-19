package main
import "fmt"

func main(){
	var name string = "Aman"
	var sname = "Singh"
	suff := "Rajput"
	fmt.Println(name , sname , suff)

	//constants grouping
	const (
		port = "5000"
		host = "localhost"
	)

	fmt.Println(host+port)
}