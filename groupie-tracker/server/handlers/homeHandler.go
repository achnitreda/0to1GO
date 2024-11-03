package handlers

import (
	"net/http"

	"web/methods"
)

// function that check for the request path and method

func HomeHandler(artists *[]methods.Artist, locations map[string][]int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			if r.URL.Path != "/" {
				ErrorHandler(w, r, http.StatusNotFound)
				return
			} else {

				data := map[string]interface{}{
					"Title":  "Home Page",
					"Artist": artists,
					"Locations": locations,
				}

				methods.RenderTemplate(w, "index.html", data)
			}

		default:
			ErrorHandler(w, r, http.StatusMethodNotAllowed)
			return
		}
	}
}
