package handlers

import (
	"net/http"

	"web/methods"
)

// function that check for the request path and method

func HomeHandler(artists []methods.Artist, locations map[string][]int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet, http.MethodPost:
			if r.URL.Path != "/" {
				ErrorHandler(w, r, http.StatusNotFound)
				return
			} else {
				// fmt.Println("artists -> ", artists, len(*artists))
				Display(w, artists, locations)
			}

		default:
			ErrorHandler(w, r, http.StatusMethodNotAllowed)
			return
		}
	}
}

func Display(w http.ResponseWriter, artists []methods.Artist, locations map[string][]int8) {
	data := map[string]interface{}{
		"Title":     "Home Page",
		"ArtistLen": len(artists),
		"Artist":    artists,
		"Locations": locations,
	}

	methods.RenderTemplate(w, "index.html", data)
}

