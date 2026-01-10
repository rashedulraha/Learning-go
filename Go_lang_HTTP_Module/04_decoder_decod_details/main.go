package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

 func writeJSON(w http.ResponseWriter ,status int  ,data any ){
	w.Header().Set("content-type" , "application/json")
	w.WriteHeader(status )
	_ = json.NewEncoder(w).Encode(data)

 }


 type TestRequest struct {
	Name string `json:"name"`

 }

//  all function here
 func testHandler ( w http.ResponseWriter ,r *http.Request){

	if r.Method != http.MethodPost{
		writeJSON(w,http.StatusMethodNotAllowed  , map[string]any{
			"ok":false,
			"error": "Only post in allowed",
		})

		return 

	}

	defer  r.Body.Close()


 	var  req  TestRequest
	dec:=json.NewDecoder(r.Body)

	if err:=dec.Decode(&req); err !=  nil {
		writeJSON(w, http.StatusBadRequest,map[string]any{
			"ok":false ,
			"error" :"Invalid json formate",
		})

		return 
	}

 }
func main() {

	http.HandleFunc("/test" ,testHandler)



	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)

}