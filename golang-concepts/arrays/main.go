package main 

import "fmt"


func main(){
	// fixed size not dynamic
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	// array literal
	res := [5]int{2,3,4,5,6};
	fmt.Println(len(res))
}