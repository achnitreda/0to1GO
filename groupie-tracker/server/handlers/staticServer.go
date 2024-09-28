package handlers

import (
	"net/http"
	"os"
)

// function that checks the request path and serve the neccessary files only

func StaticFileServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	path := "../client" + r.URL.Path
	_, err := os.Open(path)
	if err != nil {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path)
}
