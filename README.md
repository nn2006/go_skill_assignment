# Golang Skill Test
## 1.Server test:
In this Assingment I would like you to create a small go server that accepts as input a body of text, such as
that from a book, and returns the top ten most-used words along with how many times they occur in the
text.
This task should be completed using GoLang.

You can test the finctionality using following link 
http://52.31.74.51:8081/upload


## Code Walkthrough 
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
var templates = template.Must(template.ParseFiles("upload.html"))

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
	// Creating the upload route for fileuploader 
	http.HandleFunc("/upload", uploadHandler)

	//Listen on port 8081
	http.ListenAndServe(":8081", nil)
}
```
The uploadHandler() methord checks if the request is of type is GET or POST and then route  the request to the right method.

## Implementing the file uploading
Now that you have a basic HTTP server to build on let's continue by implmenting the file uploading functionality. We can do that by getting the file from the POST request we received.
```
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
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
}
```

The FieldsFunc function splits the string at each run of Unicode code points satisfying the provided function and returns an array of slices

  Read Newely created File, Each word and its frequency is stored in the wordsCount map.In order to sort the words by frequency, we create a new keys slice. We put all the words there and sort them by their frequency values. We return as response the top ten frequent words from the uploaded file
   
   
   
   
  ``` 

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
```










## Frontend
Now that the backend is ready, we need a simple frontend to act as a portal for uploading our files. For that, we will create a simple multipart form with and file input and a submit button.
```
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Upload File</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:8081/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
```
The frontend will automatically be served on the /upload endpoint when running the application.

## Testing the application
Awesome, now that we have finished the application, you can run it using the following command.
```
go run main.go
```
Or to keep this running, you can use screen or nohup

```
go build
nohup ./ServerTest &
```
