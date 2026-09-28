package main
import "fmt"

// To modify the value of x using pointers
/*func main() {               
	x := 7
	y := &x
	fmt.Println(x,y)
	*y = 8
	fmt.Println(x,y)
}*/


// Passing pointer as the parameter
func changeValue(str *string){               // The "*" on left of a datatype points to the pointer to that datatype
	*str = "changed!"           // The "*" before a varible name gets the value of the varible i.e take the pointer a figure out the value of it.
}

func changeValue2(str string){
	str = "changed!"
}

func main(){
	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)   // The "&" indicates the memory location of the varible
	fmt.Println(toChange)
}

/*func main(){
	toChange := "hello"
	var pointer *string = &toChange
	fmt.Println(pointer,&pointer)
}*/