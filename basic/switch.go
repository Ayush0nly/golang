package main 
import "fmt"

func main() {

ans := 3

/*	switch ans {
	case 1:
		fmt.Println("hello")
	case 2:
		fmt.Println("world")
	default :
		fmt.Println("hello world")
	}*/
	switch {
	case ans > 0:
		fmt.Println("hello")
	case ans < 0:
		fmt.Println("world")
	default :
		fmt.Println("hello world")
	}

}
