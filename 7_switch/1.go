package main
import "fmt"
// import "time"
func main(){
	// i := 10
	// switch i{
	// case 1:
	// 	fmt.Println("1 hai")
	// case 2:
	// 	fmt.Println("2 hai")
	// case 3:
	// 	fmt.Println("3 hai")
	// default:
	// 	fmt.Println("kuch aur hai")
	// }

	// switch time.Now().Weekday(){
	// 	case time.Saturday, time.Sunday:
	// 		fmt.Println("its weekend")
	// 	default:
	// 		fmt.Println("its weekday") 
	// }

	whoami := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("Integer")
		case bool:
			fmt.Println("Boolean")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("others", t)
		}
	}
	whoami(7)
	whoami("Aman")
	whoami(true)
	whoami(6.098)
	
}

