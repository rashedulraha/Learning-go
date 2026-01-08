package main

import (
	"fmt"
)

func main() {

	// var marks [8]int

	// marks[0] = 8

	// fmt.Println(marks)

	//! array literal

	// res := [5]int{4, 5, 2, 5, 2}

	// fmt.Println(res)

	//? common collection type
	//? dynamic and can grow
	//? []types{..........}

	// result := []string{"Rashedul" , "Nothing" ,"Everyday","Hello world"}
	// fmt.Println(result , result[0] , result[len(result)-1])
	//  result[1]="It's ok"
	//  fmt.Println(result)

	// ! append number
	var number []int
	number = append(number, 20)
	number = append(number, 50, 60 )
 	fmt.Println(number)


}