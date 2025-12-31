package main

import "fmt"

//! iterating over data structures
func main(){
	// number :=[]int{1,2,3,4,5}

	// for i:=0; i< len(number);i++{
	// 	fmt.Println(number[i])

	// }

	//?  iterating in range

	//  for _,i :=range number{
	// 	fmt.Println(i)
	//  }


	//? some in range 

			// sum:=0
			 
			// for _,i :=range number{
			// 	sum=sum + i
			// }

			//  fmt.Println(sum)



//?  map  iteration in map 

m:=map[string]string{"name":"Rashedul" , "age" : "45"}

 for k ,v :=range m {
	fmt.Println(k,v)
 }



}