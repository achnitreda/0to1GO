package handlers

import (
	"net/http"

	"web/methods"
)
// function that check for the request path and method

func HomeHandler(artists []methods.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title":  "Home Page",
			"Artist": artists,
		}

		switch r.Method {

		case http.MethodGet:
			if r.URL.Path != "/" {
				ErrorHandler(w, r, http.StatusNotFound)
				return
			} else {
				methods.RenderTemplate(w, "index.html", data)
			}

		default:
			ErrorHandler(w, r, http.StatusMethodNotAllowed)
			return
		}
	}
}
