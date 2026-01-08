package main

import (
	"errors"
	"fmt"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return  a/b , nil

}

func main() {
fmt.Println(divide(10,6))
}