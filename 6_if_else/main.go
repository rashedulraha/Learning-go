package main

import "fmt"

func main() {
	// age := 17
	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// 	}else{
	// 	fmt.Println("person is not adult")
	// }

	// else if

	age := 12

	if age >= 18 {
		fmt.Println("person in adult")
	}else if age >=12{
		fmt.Println("person in teenager")
	}else{
		fmt.Println("person in kind")
	}
}