package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{
		Name: "Aryan",
		Age:  21,
	}
	fmt.Println(u.Age)

	u.Birthday()

	fmt.Println("after" , u.Age)
}

func (u *User) Birthday() {
	u.Age++
}