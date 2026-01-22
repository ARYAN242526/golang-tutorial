package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{
		ID:    1,
		Name:  "Aryan",
		Email: "abc@google.com",
		Age:   21,
	}

	fmt.Println(u1 , u1.ID , u1.Email , u1.Name)

	// struct fields are mutable by default
	u1.Age = 25
	fmt.Println(u1);

	u2 := User {
		Name: "John",
		Email: "a@gmail.com",
	}

	fmt.Println("partial user" , u2);
}

// struct is a composite data type used to group related data of different types—similar to a class without methods in OOP languages.