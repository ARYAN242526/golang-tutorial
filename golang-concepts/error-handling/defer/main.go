package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Case 1 : success")
	if err := doWork(true); err != nil {
		fmt.Println("error:" , err)
	}

	fmt.Println("Case 2 : fail early")
	if err := doWork(false); err != nil {
		fmt.Println("error:" , err)
	}
}

func doWork(success bool) error {

	// resource related
	// start message -> resource acquired
	// cleanup message -> resource released

	fmt.Println("start : resource acquired")

	// defer will guarantee this runs at end of this function
	// both the paths
	// success return 
	// errors return - early
	

	defer fmt.Println("cleanup : resource released")

	if !success {
		return errors.New("something went wrong ! I am returning early")
	}

	fmt.Println("work : doing something imp")
	fmt.Println("work : this work is done")

	return nil
}