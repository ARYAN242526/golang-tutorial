package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "Aryan"
	lastName := "Roy"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}