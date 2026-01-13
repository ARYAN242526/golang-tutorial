package main

import "fmt"

func main(){
	// map[keyType]valueType
	ages := map[string]int{
		"aryan" : 21,
		"john" : 35,
	}

	fmt.Println(ages["aryan"] , ages["john"] , len(ages))
	

	// make(map[K]V)

	// var scores map[string]int // null map

	// fmt.Println(scores , scores["a"])

	// scores = make(map[string]int)

	// scores["math"] = 90

	// fmt.Println(scores , scores["math"])

	users := map[string]string {
		"u1" : "aryan",
		"u2" : "john",
		"u3" : "alex",
	}

	fmt.Println(users)

	delete(users , "u2")
	delete(users , "u100") // no error because u100 is not present in map

	fmt.Println(users)
}