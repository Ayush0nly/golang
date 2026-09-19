package main

import "fmt"


/*func test (x int) {
	fmt.Println("hello!", x)
}

func main() {
	x := test                            // Assigning function to a varible
	x(5)
}*/

/* Creating a function inside a fuction
func main(){
	a := func(y int) int {
		return y*-1
	}(10)
	fmt.Println(a)
}*/

// Passing functions as parameters
/*func test2(MyFunc func(int) int ){
	fmt.Println(MyFunc(7))
}

func main() {
	test := func(x int) int{
		return x*-1
	}

	test3 := func(x int) int{
		return x * 7
	}
	test(5)
	test2(test3)
}*/

// Fuction closure
func returnFunc(x string) func(){
	sum := 0
	return func(){                       // Function closure
		fmt.Println(sum)
	}
}

func main () {
	returnFunc("hello!")()
	x := returnFunc("goodbye")
	x()
}