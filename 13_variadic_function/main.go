package main

import "fmt"

func sum(number ...int) int {

	total := 0

	for _, num := range number {
		total = total + num
	}

	return total

}

func main() {

	result := sum(2, 4, 5, 4, 54, 5, 4, 5, 5 , 45)
	fmt.Println(result)

}