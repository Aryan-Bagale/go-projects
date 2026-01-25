package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AgeRequest struct {
	Age int `json:"age"`
}

func main() {
   
	http.HandleFunc("/enter", func(w http.ResponseWriter, r *http.Request) {
        
        var data AgeRequest
        
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest) 
			fmt.Fprint(w, "Invalid JSON data")
			return 
		}
        
		if data.Age < 18 {
			w.WriteHeader(http.StatusUnauthorized) // 401
			fmt.Fprint(w, "You are too young")
		} else {
			w.WriteHeader(http.StatusOK) // 200
			fmt.Fprint(w, "Come in!")
		}
	})

	http.ListenAndServe(":8080", nil)
}