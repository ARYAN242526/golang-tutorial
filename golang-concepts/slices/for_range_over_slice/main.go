package main

import "fmt"

func main(){
	views := []int{10,20,30,40,50}

	// for range
	total := 0
	for i , val := range views {
		fmt.Println("day" , i , "views" , val)
		total = total + val
	}

	fmt.Println(total)

}