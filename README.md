# Golang Skill Test
## 1.Server test:
In this Assingment I would like you to create a small go server that accepts as input a body of text, such as
that from a book, and returns the top ten most-used words along with how many times they occur in the
text.
This task should be completed using GoLang.

You can test the finctionality using following link 
http://52.31.74.51:8081/upload


#### Code Walkthrough 
The server having single endpoint with a GET and POST request implementation. Initially the GET request should display the frontend. A POST request, on the other hand, should trigger the file uploading process.
```
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

// Compile templates on start of the application
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func main() {
	// Upload route
	http.HandleFunc("/upload", uploadHandler)

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
```
