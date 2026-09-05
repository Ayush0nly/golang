package main 
import "fmt"

func main() {
	/*x:=2
	for x<=5 {
		fmt.Println(x)
		x++
	} multi line comment */

	// for x:=1; x<=5; x+=2 {
	// 	fmt.Println(x)
	// }another way of writing for loops

	//for loop for continue and break
	for x:=1; x<=1000; x++ {
		if x!=1 && x%3==0 && x%7==0 && x%9==0{
			fmt.Println(x)
			// continue loops works until the maximum condition is reached
			// break only first output then break out of the loop
		}
	}
}