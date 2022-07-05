package handlerFunctions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	blogLogic "simpleCollabBlog/simple-blog-backend/blogLogic"
	"simpleCollabBlog/simple-blog-backend/dataStructs"
	"simpleCollabBlog/simple-blog-backend/testData"
	"time"

	"github.com/gorilla/mux"
)

//-----------------------------------------------------------------------------
/// Function to send back a test struct. This function is only for testing purpose. Should be deleted at some point
func SendTestData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Print("HELLO")
	json.NewEncoder(w).Encode(testData.TestData)
}

//-----------------------------------------------------------------------------
/// Function to read the body from a POST request and marshall the content of the post to a dummy struct
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

//-----------------------------------------------------------------------------
/// Function to transform a Mardown content to a HTML content
/// The content of the Markdown is in the Body of the POST request
func SendMdToHTML(w http.ResponseWriter, r *http.Request) {
	var req dataStructs.BlogStructure

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	blogStructureHTML := blogLogic.MdStructToHTML(req)
	var blogHTMLstruct dataStructs.BlogHTML
	blogHTMLstruct.HtmlContent = string(blogStructureHTML)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogHTMLstruct)
}

//-----------------------------------------------------------------------------
/// Function to save a Markdown file to the local filesystem
func SaveMd(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	var req dataStructs.BlogStructure
	err := json.NewDecoder(r.Body).Decode(&req)

	err = os.WriteFile("blogArticles/"+req.Title+" "+currentTime.Format("2006-January-02")+".md", []byte(req.Content), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

//-----------------------------------------------------------------------------
/// Function to upload an existing Markdown
func UploadMarkdownFile(w http.ResponseWriter, r *http.Request) {
	var blog dataStructs.BlogStructure
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	filePath := "blogArticles/" + fileName + ".md"
	file, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	blog.Content = string(file)
	blog.Title = "None"

	blogStructureHTML := blogLogic.MdStructToHTML(blog)
	var blogHTMLstruct dataStructs.BlogHTML
	blogHTMLstruct.HtmlContent = string(blogStructureHTML)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogHTMLstruct)

}
