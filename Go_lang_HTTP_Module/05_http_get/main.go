package main

import (
	"fmt"
	"net/http"
)

func main() {
	// json placeholder
	url := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return 
	}

	defer res.Body.Close()
	fmt.Println("status code:",res.StatusCode)
	fmt.Println("status:",res.Status)

}