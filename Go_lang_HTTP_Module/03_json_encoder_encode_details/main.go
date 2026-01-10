package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//

//!  json encoder encode details

 func successHandler (w http.ResponseWriter ,  r *http.Request){

	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(http.StatusOK)

	res:= map[string]any{
		"ok":true ,
		"message":"json encoded successfully and all complete",
		"dateTime":time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)


 }


func main() {

	http.HandleFunc("/ok" , successHandler)

	err:= http.ListenAndServe(":8080",nil)
	fmt.Println(err)

}