package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"web/handlers"
)

func staticFileServer(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		handlers.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	path := "../client" + r.URL.Path
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		handlers.ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path)
}

func main() {
	http.HandleFunc("/static/", staticFileServer)

	http.HandleFunc("/download", handleDownload)

	http.HandleFunc("/", handlers.Homehandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	fileType := r.URL.Query().Get("type")
	content := r.URL.Query().Get("content")

	var fileName string
	var contentType string

	if fileType == "text" {
		fileName = "file.txt"
		contentType = "text/plain"
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	w.Header().Set("Content-Type", contentType)
	w.Write([]byte(content))
}
