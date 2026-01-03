package main

import "fmt"

// print slice

  func printSlice[T int | string | bool](items []T){
		for _,item :=range items{
			fmt.Println(item)
		}
	}
  // func printStringSlice(items []string){
	// 	for _,item :=range items{
	// 		fmt.Println(item)
	// 	}
	// }

func main() {
	// fmt.Println("Hello Generics")

	// number:=[]int{1,2,3,4,5,6,7,8,9}
	// printSlice(number)
	// PersonName :=[]string{"Rashedul Islam" , "Abdullah" ,"Rabbi","Nothing"}
	
	values:=[]bool{true , false }
	printSlice(values)



}