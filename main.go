package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Defines a struct for data
type Data struct {
	Value int `json:"value"`
}

// Defines a struct for analysis result
type AnalysisResult struct {
	Average float64 `json:"average"`
}

var dataStore []Data

func submitData(w http.ResponseWriter, r *http.Request) {
	var newData []Data
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appends the array of values to the dataStore
	dataStore = append(dataStore, newData...)

	w.WriteHeader(http.StatusCreated)
}


func getAnalysisResults(w http.ResponseWriter, r *http.Request) {
	if len(dataStore) == 0 {
		http.Error(w, "No data for analysis", http.StatusNotFound)
		return
	}

	// Simple average calculation for API demonstration
	var sum int
	for _, data := range dataStore {
		sum += data.Value
	}
	average := float64(sum) / float64(len(dataStore))

	result := AnalysisResult{Average: average}
	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()

	// Defines API endpoints
	r.HandleFunc("/data", submitData).Methods("POST")
	r.HandleFunc("/analysis", getAnalysisResults).Methods("GET")

	// Starts the server
	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
