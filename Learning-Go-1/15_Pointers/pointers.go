package main

import "fmt"

// func changeNum(num int) {
// 	num = 5

// fmt.Println("In change Number", num)
// }

// by reference

func changeNum(number *int){
	*number = 45
	fmt.Println("In changeNum" , *number)

}


func main() {


	num:=57
	changeNum(&num)
	fmt.Println("memory address" , num)
	// fmt.Println("After change in main function" , num)
}