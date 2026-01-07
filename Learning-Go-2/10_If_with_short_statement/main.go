package main

import "fmt"

func main() {
	item := 3

	pricePerItem := 45

	if total := item * pricePerItem; total >= 100 {
		fmt.Println("Eligible for shipping")
	}else{
		fmt.Println("not eligible shipping")
	}

}