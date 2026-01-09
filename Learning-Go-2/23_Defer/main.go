package main

import (
	"errors"
	"fmt"
)

func main() {

	// Defer  resp.body.close()

	// defer fmt.Println("1")
	// fmt.Println("2")
	// fmt.Println("3")

	fmt.Println("Case 1: success")
	if err:= doWork(true);err !=nil{
	fmt.Println("Error:" , err)
	}

	fmt.Println("Case 1: false")
	if err:= doWork(false);err !=nil{
	fmt.Println("Error:" , err)
	}

}

func doWork(success bool) error {
	fmt.Println("Start:","resource acquired" )

	defer fmt.Println("cleanup : resource released")

	if !success {
		return  errors.New("Something went wrong , i'm returning early")
	}

	fmt.Println("doing something imp")
	fmt.Println("This work is done")

	return  nil
	
}