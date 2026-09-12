package main
import "fmt"

func main(){
	arr := [] int {3, 4, 5}     // to sum the element of array
	sum := 0

	for i:=0; i<len(arr); i++{
		sum += arr[i]
	}
	fmt.Println(sum)

	arr2D := [2][3] int {{1,2,3},{4,5,6}}  // to create a array inside array
	fmt.Println(arr2D[1][1])

	// var arr [5]int              another way to create array

	// fmt.Println(len(arr))        to find the length of array

	// arr[0]=100             to change the element at 0th index

	// fmt.Println(arr[0])       to print the element at the 0th index

}
