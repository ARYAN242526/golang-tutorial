package main

import (
	"fmt"
)

func main(){
	isLogged := true; // inferred as boolean type
	isAdmin := true;
	hasSubscription := false;

	// AND && 
	canOpenDashBoard := isLogged && hasSubscription

	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println(canOpenDashBoard , canDeletePost)

	age := 24
	isAdult := age > 18
	fmt.Println(isAdult)
}