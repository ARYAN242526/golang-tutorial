package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sumAndProduct(a int , b int) (int , int) {
	sum := a + b
	product := a * b

	return sum , product
}

func main() {
	sum1 := add(10, 20)

	s, p := sumAndProduct(6 , 5)
	fmt.Println(sum1 , s , p)

	// _ blank identifier used to ignore values
	onlySum , _ := sumAndProduct(10 , 2)
	fmt.Println(onlySum)

}