package main

import (
	"fmt"
	"go-modules/Internal/greet"
)

func main() {
msg1:=greet.Hello("Rashedul")

fmt.Println(msg1)
}