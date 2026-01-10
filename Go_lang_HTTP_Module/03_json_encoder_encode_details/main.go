package main

import (
	"fmt"
	"net/http"
)

//!  json encoder encode details

 func successHandler (w http.ResponseWriter ,  r *http.Request){

	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(http.StatusOK)


 }


func main() {

	http.HandleFunc("/ok" , successHandler)

	err:= http.ListenAndServe(":8080",nil)
	fmt.Println(err)

}