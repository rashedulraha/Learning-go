package main

import "fmt"

// slices -> dynamic
// most use constructs in  go
// + useful method

 func main(){

	// uninitialized slices is nil

	// var name []int
	// fmt.Println(name==nil)

	// var number =make([]int, 6 ,50)
	// fmt.Println(cap(number))


	// number = append(number, 5)
	// fmt.Println(cap(number))
	// fmt.Println(number)


	//  fmt.Println(number == nil)


	//! copy function

	 var number = make([]int, 2, 5)

	  number[0]=4

	 fmt.Println(number)

 }