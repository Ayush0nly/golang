package main

import "fmt"

// Making our own type using structs
type Point struct {
	x int32
	y int32
	str string
}

func main() {
	p1 := Point{x:1,}
	var p2 Point = Point{-5, 7, "Raj"}
	var p3 Point = Point{2, 6, "Yadav"}
	p1.x = 3                      // To change the value of point at x
	fmt.Println(p1)
	fmt.Println(p2.y)
	fmt.Println(p3.str)
}