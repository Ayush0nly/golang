// arithmetic operations
package main
import (
	"fmt"
)

func main() {
	var num1 int = 5
	var num2 int = 10
	answer := num1 - num2
	fmt.Printf("%d", answer)

// conditions & boolean expression

	x:= 10   //we can compare string x:="a" and y:="b"
	y:= 5.5
	// val:= float64(x)+1.5 != y
	fmt.Printf("%t",float64(x)+1.5 != y)
	// different ways to write  

// chained conditionals(AND,OR,NOT!)

	z := 5
	val:= (true || false) && !false ||z>9
	val1:= val || false
	fmt.Printf("%t", val1)
}