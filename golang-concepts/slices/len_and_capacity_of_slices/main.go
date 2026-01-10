package main

import "fmt"

func main(){
	// make([]type , len , cap)

	// scores := make([]int,0,5)

	// fmt.Println(scores , len(scores) , cap(scores))

	// scores = append(scores, 100)
	// fmt.Println("after appending 100" , scores)

	// scores = append(scores, 200 , 500)
	// fmt.Println("after appening 200 , 500" , scores)

	// scores = append(scores, 45 , 55)
	// fmt.Println("agter appending 45 55" , scores)

	// // in case if we are exceeding capacity , go usually doubles the capacity of slice
	// scores = append(scores, 60)
	// fmt.Println("after appending 60" , scores , len(scores) , cap(scores))

	todos := []string{"make vlogs" , "workout"}

	more := []string{"learn go"}

	//... spreads one slice into another 
	todos = append(todos, more...)
	fmt.Println(todos)
}