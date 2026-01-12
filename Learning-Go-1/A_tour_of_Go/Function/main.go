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
//  func swap (x  , y string)(string ,  string){
// 	return  x, y

//  }

//! Named return values function

 func nakedReturn  ( name string , age int , address string)(PersonName string , personAge string , personAddress string ){
	PersonName= "Hello" + name
	personAge = fmt.Sprintf("person age is %d" , age)
	personAddress= "Your address is " + address
	 return 
 }





func main() {
	// fmt.Println(add(5,56))
	//  a,b := swap("Hello" , "World")
	// 	fmt.Println(a,b)


	fmt.Println(nakedReturn("Rashedul" , 21 , "Naogaon"))
}