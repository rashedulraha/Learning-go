package main

import "fmt"

// print slice

  func printSlice(items []int){
		for _,item :=range items{
			fmt.Println(item)
		}
	}

func main() {
	// fmt.Println("Hello Generics")

	printSlice([]int{1,2,3,4,5,6,7,8,9})
}