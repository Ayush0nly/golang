package main 
import "fmt"

type Student struct{
	name string
	grades []int
	age int
}

/*func (s *Student) setAge(age int){
	s.age = age
}*/

// Calculating average grade
func (s *Student) getAverageGrade() float32{
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

// Finding maximum grade
func (s *Student) getMaxGrade() int {
	curMax := 0

	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}

	return curMax
}

func main(){
	s1 := Student{"john", []int{70,80,90}, 21}
	s2 := Student{"tim", []int{75,90,95}, 19}
	
	// fmt.Println(s1)
	// s1.setAge(7)
	// fmt.Println(s1)

	// average := s1.getAverageGrade()
	// average2 := s2.getAverageGrade()
	// fmt.Println(average,average2)

	max := s1.getMaxGrade()
	max2 := s2.getMaxGrade()
	fmt.Println(max, max2)
}