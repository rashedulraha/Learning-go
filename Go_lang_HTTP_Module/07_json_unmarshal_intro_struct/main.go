package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type catFactResponsiveStructures struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)

	if err != nil{
		fmt.Println(err)
		return 
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println(resp.Status)
		return 
	}


	bodyByte,err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return 
		
	}

	var data catFactResponsiveStructures

	if err :=  json.Unmarshal(bodyByte, &data); err != nil {
		fmt.Println("json unmarshal failed")
		return
	}

	fmt.Print(data.Fact, data.Length)



}