package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

//!  shortened Function
// func add (x , y int)int{
// 	return x+y
// }

//! multiple result
 func swap (x  , y string)(string ,  string){
	return  x, y

 }





func main() {
	// fmt.Println(add(5,56))
	 a,b := swap("Hello" , "World")
		fmt.Println(a,b)
}