package main

import "fmt"

func main() {
	//! simple switch

	i := 6

	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	case 4:
		println("four")
	case 5:
		println("five")

	default:
		println("no match here")

	}

	//!  multiple condition switch

	// switch time.Now().Weekday(){

	// case time.Saturday,time.Tuesday:
	// 	println("it's weekend")

	// 	 default :
	// 	 println("it's workday")

	// }

	// ! type switch

	whoAmi := func(i any) {
		switch t := i.(type) {
		case int:
			fmt.Println("ist an integer")
		case string :
			fmt.Println("i'st string")
		case bool :
			fmt.Println("i's boolean")
		default: 
		fmt.Println("other" , t)
		}

	}

	
whoAmi(45.45)

}


