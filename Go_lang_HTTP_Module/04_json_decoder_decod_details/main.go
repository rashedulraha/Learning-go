package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//? start helper function
 func writeJSON(w http.ResponseWriter ,status int  ,data any ){
	// set response header
	w.Header().Set("content-type" , "application/json")
	//  add status code
	w.WriteHeader(status )
	// convert json formate 
	_ = json.NewEncoder(w).Encode(data)

 }

//? end helper function


 type TestRequest struct {
	Name string `json:"name"`
 }

//  all function here

//  here 5 important work  in function
 func testHandler (w http.ResponseWriter ,r *http.Request){


	// method filtering 
	if r.Method != http.MethodPost{
		writeJSON(w,http.StatusMethodNotAllowed  , map[string]any{
			"ok":false,
			"error": "Only post in allowed",
		})
		return 
	}

	// work finished to close pipeline using defer
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

	req.Name= strings.TrimSpace(req.Name)
 
	if req.Name==""{
		writeJSON(w, http.StatusBadRequest,map[string]any{
			"ok":false,
			"error":"Name must not be empty",
		})
		return 
	}
	writeJSON(w, http.StatusOK,  map[string]any{
		"ok":true, 
		"data":req, 
		"timeStamp":time.Now().UTC(),
	})

 }
func main() {

	http.HandleFunc("/test" ,testHandler)



	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)

}