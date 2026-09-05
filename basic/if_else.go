package main

import "fmt"

func main() {
	age := 17
	if age >= 18 {
		fmt.Println("you can ride!")
	} else if age >= 14 {
		fmt.Println("you can ride with parent!")
	} else {
		fmt.Println("you cannot ride!")
	}
}
