package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0

	for _, currentVal := range nums {
		total += currentVal
	}

	return total
}

func main() {
	fmt.Println(sumAll(1,2,3,4,5,6,7,8))

	values := []int{10,20,30}
	fmt.Println(sumAll(values...))

	res := func(n int) int{
		return n * 2
	}

	fmt.Println(res(2))

	// IIFE

	res1 := func (a int , b int) int {
		return a - b
	}(10 , 5)

	fmt.Println(res1)

}

// Variadic function -> function that accepts a variable number of arguments of the same type
