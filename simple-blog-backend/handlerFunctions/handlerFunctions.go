package handlerFunctions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleCollabBlog/simple-blog-backend/dataStructs"
	"simpleCollabBlog/simple-blog-backend/testData"
)

func SendTestData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testData.TestData)
}

func GetTestData(w http.ResponseWriter, r *http.Request) {
	var req dataStructs.TestDataStruct

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	structString, _ := json.MarshalIndent(req, "", "\t")
	fmt.Fprintf(w, string(structString)+"\n")
	fmt.Fprintf(w, req.Title+"\n")
	fmt.Fprintf(w, req.Content+"\n")
}
