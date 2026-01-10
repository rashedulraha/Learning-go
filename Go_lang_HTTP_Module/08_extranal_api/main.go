package main

import (
	"fmt"
	"net/http"
)

type catFactResponse struct  {
	Fact string `json:"fact"`
	Length int `json:"length"`

}



 func externalHandler ( w http.ResponseWriter , r *http.Request){

 }

func main() {

	http.HandleFunc("/external" , externalHandler)

	err := http.ListenAndServe(":5300", nil)
	fmt.Println(err)

}