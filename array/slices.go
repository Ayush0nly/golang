package main
import "fmt"

func main(){
	var a []int = []int{5, 10, 15, 20, 25}

	a = append(a, 30)      // add element to slice
	fmt.Println(a)

	fmt.Println(len(a[:3]))

	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int =x[1:4]                  // to find the length of slice use len(s) & for capacity use cap(s)

	fmt.Println(s[:cap(s)])              // slice of a slice (s[1:3]) ,, len(s) goes till length ,, cap(s) goes till last element

	b := make([]int,5)                   // slice using make keyword
	fmt.Printf("%T",b)
}
