package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("Upload template")
		t, _ = t.Parse(uploadTemplate)
		t.Execute(w, nil)
	})

	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}

var uploadTemplate = `
<!DOCTYPE html>
<html>
<body>
    <h2>CSV File Upload</h2>
    <form enctype="multipart/form-data" action="/upload" method="post">
        <input type="file" name="file" accept=".csv">
        <input type="submit" value="Upload">
    </form>
</body>
</html>
`

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		fmt.Println(record)
		fmt.Fprintln(w, record)
	}
}
