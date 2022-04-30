package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"log"
	"sort"
	"strings"
    "io/ioutil"   
)

// Compile templates on start of the application
var templates = template.Must(template.ParseFiles("upload.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
//	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


   //Read Newely created File 
   
   

	filecontant, err := ioutil.ReadFile(handler.Filename)

	if err != nil {

		log.Fatal(err)
	}

	text := string(filecontant)

	fields := strings.FieldsFunc(text, func(r rune) bool {

		return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
	})

	iwordsCount := make(map[string]int)

	for _, field := range fields {

		iwordsCount[field]++
	}

	keys := make([]string, 0, len(iwordsCount))

	for key := range iwordsCount {

		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {

		return iwordsCount[keys[i]] > iwordsCount[keys[j]]
	})

	for idx, key := range keys {

          fmt.Fprintf(w, "%s %d\n", key, iwordsCount[key])
		

		if idx == 10 {
			break
		}
	}

	
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

	//Listen on port 8081
	http.ListenAndServe(":8081", nil)
}