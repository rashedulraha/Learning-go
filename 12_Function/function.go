package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

//  string  function


 func stringValue ()(string,string,string){
	return "Javascript" , "typescript", "next.js"
 }

//?  feature
func add(a, b int) int {
	return a + b
}

func main() {

	// result := add(5, 4)
	// fmt.Println(result)

	//?  call string value function


	 valueOne , valueTwo , valueThree  := stringValue()
	 fmt.Println(valueOne , valueTwo, valueThree )

	//  fmt.Println(stringValue())

}