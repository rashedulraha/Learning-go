package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//

//!  json encoder encode details

//? start success handler function
 func successHandler (w http.ResponseWriter ,  r *http.Request){

	// dec content type json formate 
	w.Header().Set("content-type" , "application/json")

	// confirm browser request all okey
	w.WriteHeader(http.StatusOK)


	// create a go map  including dynamic 3 kay
	res:= map[string]any{
		"ok":true ,
		"message":"json encoded successfully and all complete",
		"dateTime":time.Now().UTC(),
	}

	// convert json and send data to browser
	_ = json.NewEncoder(w).Encode(res)


 }

 //? End success handler function
 


func main() {
// hit path to run successHandler function
	http.HandleFunc("/ok" , successHandler)


	// run server port number 8080
	err:= http.ListenAndServe(":8080",nil)

	// print err
	fmt.Println(err)

}