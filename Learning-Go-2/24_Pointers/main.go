package main

import "fmt"

func main() {

	// age := 25

	// var p *int

	// p = &age

	// fmt.Println("Age", age)
	// fmt.Println("Age", p)
	// fmt.Println("Address", *p)

	// ! store the memory address of any value

	// &x -> address  of x (make a pointer)
	// *p -> defer ( go that address and read/write)

	score := 10
	fmt.Println("before:",score)

	addScore(score)
}

func addScore(score int) {
fmt.Println(score)
}