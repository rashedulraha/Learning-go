package main

import (
	"fmt"
	"io"
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

 if res.StatusCode != http.StatusOK{
	fmt.Println(res.Status)
 }



 bodyBytes, err := io.ReadAll(res.Body)

 	if err != nil {
		fmt.Println(err)
		return
	}

	bodyText:= string(bodyBytes)

	max:=205

	if (len(bodyText) <max){
		max =len(bodyText)
	}

	fmt.Println(bodyText[:max])
	
}

