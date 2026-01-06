package main

import "fmt"

func main() {

	score := 30

	if score >= 90 {
		fmt.Println("A")
	}else if score >= 75{
		fmt.Println("B")
	} else if  score  >= 45{
		fmt.Println("C")
	}else{
		fmt.Println("D")
	}

}