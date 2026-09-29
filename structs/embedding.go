// Struct under a struct (ex of a circle and centre)
package main 
import "fmt"
 
type Point struct{
	x int32
	y int32
}

type Circle struct{
	radius float32
	centre *Point
}

func main (){
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1.centre.x)
}