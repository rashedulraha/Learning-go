package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most use constructs in  go
// + useful method

 func main(){

	// uninitialized slices is nil

	var name []int
	fmt.Println(name==nil)

	var number =make([]int, 6 ,50)
	fmt.Println(cap(number))


	number = append(number, 5)
	fmt.Println(cap(number))
	fmt.Println(number)


	 fmt.Println(number == nil)


	//! copy function

	//  var number = make([]int, 0, 5)
	 number =append(number, 1)
	//  var number2  =make([]int, len(number))


	//  copy(number2,number)

	//  fmt.Println(number,number2)



	//  ! slice operator 

	// var number   =[]int{1,2,3,4,5}
	// fmt.Println(number[2:4])
	// fmt.Println(number[2])
	// fmt.Println(number[:2])
	fmt.Println(number[1:])



	//  slices in build feature 
	 var number1 = []int{1,2, 5}
	 var number2  = []int{1,2 , 4}


	 fmt.Println(slices.Equal(number1,number2))



	//  slice 2d array 

	 var array  =[][]int{{3,4,5},{6,7,8}}

	 fmt.Println(array)

 
	






  
 }