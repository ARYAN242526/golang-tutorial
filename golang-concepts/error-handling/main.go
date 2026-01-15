package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "3"

	level , err := parseLevel(input)
	if err != nil {
		return err
	}

	fmt.Println("selected level" , level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value , err)
	// nil error -> success
	// not nil -> failure

	n , err := strconv.Atoi(s)
	if err != nil {
		return 0 , fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0 , fmt.Errorf("level must be between 1 and 5")
	}

	return  n , nil
}

// go doesnt use exception for normal failures
// functions -> return errors as normal return values

// val , err := something()
// if err != nil (handle the error)