package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type testDataStruct struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var testData = testDataStruct{
	Title:   "SAND",
	Content: "HASS",
}

func main() {

	r := mux.NewRouter()

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swaggerui/").Handler(sh)

	r.HandleFunc("/api/test", sendTestData).Methods("GET")
	r.HandleFunc("/api/rec", getTestData)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func sendTestData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testData)
}

func getTestData(w http.ResponseWriter, r *http.Request) {
	var req testDataStruct

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	structString, _ := json.MarshalIndent(req, "", "\t")
	fmt.Fprintf(w, string(structString)+"\n")
	fmt.Fprintf(w, req.Title+"\n")
	fmt.Fprintf(w, req.Content+"\n")
}
