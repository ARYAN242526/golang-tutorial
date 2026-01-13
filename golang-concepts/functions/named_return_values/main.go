package main

import "fmt"

func divide(a int, b int) (x int, y int) {
	x = a / b
	y = a % b

	return
}


func main() {

	q, r := divide(10, 10)
	fmt.Println(q,r)
}