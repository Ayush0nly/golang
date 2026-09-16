package main
import "fmt"
func main() {
	var a []int = []int{1,2,2,3,3,4,5,5,6,7}             // To print the elements of a slice using for loop
	for i:=0; i<len(a);i++{
		fmt.Println(a[i])
	}

	for i, element := range a {                   // To loop through all the element of array using range keyword
		fmt.Printf("%d:%d\n",i,element)           // We can use _ in place of i oe element as a anonymous 
	}

// To print out the duplicate in the array
	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && j > i {
				fmt.Println(element)
			}
		}
	}
// To print out duplicate other way
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
