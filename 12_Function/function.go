package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {

	result := add(5, 5)

	fmt.Println(result)
}