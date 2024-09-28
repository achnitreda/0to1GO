package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"web/methods"
)
// function that checks for the request path and method
// also calls the FetchParser

func ArtistHandler(artists []methods.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if strings.HasPrefix(path, "/artists/id=") {
			id := strings.TrimPrefix(path, "/artists/id=")
			nId, err := strconv.Atoi(id)
			if err != nil || nId <= 0 || nId > len(artists) {
				ErrorHandler(w, r, http.StatusNotFound)
				return
			}

			switch r.Method {
			case http.MethodGet:
				var relation methods.Relation
				var locations methods.Locations
				var dates methods.Dates
				err = methods.FetchParser(artists[nId-1].Relations, &relation)
				if err != nil {
					ErrorHandler(w, r, http.StatusInternalServerError)
					return
				}

				err = methods.FetchParser(artists[nId-1].Locations, &locations)
				if err != nil {
					ErrorHandler(w, r, http.StatusInternalServerError)
					return
				}

				err = methods.FetchParser(artists[nId-1].ConcertDates, &dates)
				if err != nil {
					ErrorHandler(w, r, http.StatusInternalServerError)
					return
				}
				data := map[string]interface{}{
					"Title":    "Artist Page",
					"Artist":   artists[nId-1],
					"Relation": relation,
					"Dates": dates,
					"Locations": locations,
				}
				methods.RenderTemplate(w, "artist.html", data)
			default:
				ErrorHandler(w, r, http.StatusMethodNotAllowed)
				return
			}

		} else {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}
	}
}
