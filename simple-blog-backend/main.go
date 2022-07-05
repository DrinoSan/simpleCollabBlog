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

	// Markdown Endpoints

	// Enpoint to transform a markdown file to a html output
	r.HandleFunc("/api/mdtohtml", hfunc.SendMdToHTML).Methods("POST")
	r.HandleFunc("/api/savemd", hfunc.SaveMd).Methods("POST")
	r.HandleFunc("/api/uploadmd/{fileName}", hfunc.UploadMarkdownFile).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
