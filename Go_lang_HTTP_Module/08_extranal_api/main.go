package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//   data structure model
type catFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// helper function to send data
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// 3 to get data from external api

func fetchCatFact() (catFactResponse, error) {
	url := "https://catfact.ninja/fact"
	resp, err := http.Get(url)

	if err != nil {
		return catFactResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return catFactResponse{}, fmt.Errorf("external api failed with status: %s", resp.Status)
	}


	//  get data  to convert decoded and insert struct
	var data catFactResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	
	if err != nil {
		return catFactResponse{}, err
	}

	return data, nil
}


//   call external handler  
func externalHandler(w http.ResponseWriter, r *http.Request) {

	fact, err := fetchCatFact()

	if err != nil {

		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":   true,
		"data": fact,
	})
}

func main() {

	http.HandleFunc("/external", externalHandler)

	fmt.Println("Server is running on http://localhost:5300")

	err := http.ListenAndServe(":5300", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}