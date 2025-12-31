package main

import "fmt"

func changeNum(num int) {
	num = 5

	fmt.Println("In change Number", num)

}

func main() {


	num:=1
	changeNum(num)
	fmt.Println("After change in main function" , num)

}