package main
import "fmt"
func main() {
	var fruits map[string]int = map[string]int{
		"Apple":5,
		"pears":10,
		"orange":15,
	}

// Other way to make map
	emptymp := make(map[string]int)
	fmt.Println(emptymp)

	fmt.Println(fruits["Apple"])       // To access Apple            // Maps store data in key value pair in unshorted way

	fruits["mango"]=0                  //To add new key value or to change value of a key
	delete(fruits, "pears")

	val, ok := fruits["orange"]       //To check if the value exist
	fmt.Println(val,ok) 

	fmt.Println(fruits)
}