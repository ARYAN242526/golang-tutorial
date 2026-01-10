package main

import "fmt"

func main(){
	// slices -> dyanmic size
	// []type{...}

	results := []string{"aryan" , "john"}
	fmt.Println(results , results[0] , results[len(results) - 1])

	results[1] = "alex"
	fmt.Println(results)

	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20,30)

	fmt.Println(nums)
}