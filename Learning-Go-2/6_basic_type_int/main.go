package main

import "fmt"

func main() {

	views1 := 5000
	views2 := 8000
	totalViews := views1 + views2
	avgViews := totalViews/2

	like := 50
	like++
	like++

	fmt.Println(totalViews, like , avgViews)

}