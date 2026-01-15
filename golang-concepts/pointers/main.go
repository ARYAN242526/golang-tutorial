package main

import "fmt"

func main() {
	// store memory address
	// &x -> address of x (Address operator)
	// *p -> dereference (access the value at particular memory address)

	score := 10
	fmt.Println("before: ", score)

	addScore(&score)
	fmt.Println("after: " , score)
}

func addScore(score *int) {
	*score = *score + 5
}