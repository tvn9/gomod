package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tvn9/gomod/toolkit"
)

func main() {
	mux := routes()

	log.Println("Starting server on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("error loading server", err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/upload", uploadFiles)
	mux.HandleFunc("/upload-one", uploadOneFile)

	return mux
}

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method  not allowed", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      1020 * 1024 * 1024,
		AllowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
	}

	files, err := t.UploadFile(r, "./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	out := ""

	for _, i := range files {
		out += fmt.Sprintf("Uploaded %s to the uploads folder, renamed to %s", i.OriginalFileName, i.NewFileName)
	}

	_, _ = w.Write([]byte(out))
}

func uploadOneFile(w http.ResponseWriter, r *http.Request) {

}
