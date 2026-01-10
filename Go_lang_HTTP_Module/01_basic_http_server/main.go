package main

import (
	"fmt"
	"net/http"
)


func handleHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w , "only allow get request" , http.StatusMethodNotAllowed)
		return 
	}

	_,_ = w.Write([]byte("Hello from go net/http"))

}

func main() {

http.HandleFunc("/hello" , handleHandler)

fmt.Println("try going to 8000 port")
err:= http.ListenAndServe(":8000" , nil)
fmt.Println(err)

}