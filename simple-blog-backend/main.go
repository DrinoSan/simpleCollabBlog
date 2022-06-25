package main

import (
	"log"
	"net/http"
	hfunc "simpleCollabBlog/simple-blog-backend/handlerFunctions"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swaggerui/").Handler(sh)

	r.HandleFunc("/api/test", hfunc.SendTestData).Methods("GET")
	r.HandleFunc("/api/rec", hfunc.GetTestData).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
