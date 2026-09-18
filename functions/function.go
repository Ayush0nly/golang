package main
import "fmt"
 
func test() {                   // Creating a test function
	fmt.Println("test")
}

func add(x int, y int) {        //parameters of the function add which are x and y  (x, y int)
	fmt.Println(x + y)
}

// Storing the value tins return varible
func eq(x, y int) (z1 int , z2 int) {
	defer fmt.Println("hello")                  // defer function print at the end of the return varible

/* labelling the return varible
	z1 = x + y
	z2 = x - y */

	fmt.Println("before return")
	return x + y , x - y                       //Multiple return varibles 
}

func main() {
	test()                     // Calling the test function will print whatever is inside it

	ans1 , ans2 := eq(3 , 4)   // Argument of eq function
	fmt.Println(ans1 , ans2)

	add(5,6)                    // Arguments of the function add
}
