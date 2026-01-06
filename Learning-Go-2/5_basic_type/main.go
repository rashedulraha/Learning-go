package main

import (
	"fmt"
	"strings"
)

func main() {

	firstName := "rashedul"
	lastName := "islam"

	fullName := firstName + " " + lastName
	//  convert to toUpper

	 toUpper:= strings.ToUpper(fullName)

	fmt.Println(toUpper)

}