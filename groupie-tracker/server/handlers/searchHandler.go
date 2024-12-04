package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"web/methods"
)

type SearchSuggestion struct {
	ID    int
	Title string
	Type  string
}

func SearchHandler(artists []methods.Artist, locations map[string][]int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			query := r.URL.Query().Get("q")

			var res []SearchSuggestion

			if query == "" {
				json.NewEncoder(w).Encode(artists)
				return
			}

			query = strings.ToLower(query)
			for _, artist := range artists {
				if strings.Contains(strings.ToLower(artist.Name), query) {
					// continue to skip (phill collins has one member is himself)
					res = append(res, SearchSuggestion{
						ID:    artist.Id,
						Title: artist.Name,
						Type:  "- artist/band",
					})
					continue
				}

				creationDate := strconv.Itoa(artist.CreationDate)
				if strings.Contains(creationDate, query) {
					res = append(res, SearchSuggestion{
						ID:    artist.Id,
						Title: creationDate,
						Type:  "- Creation Date",
					})
				}

				if strings.Contains(artist.FirstAlbum, query) {
					res = append(res, SearchSuggestion{
						ID:    artist.Id,
						Title: artist.FirstAlbum,
						Type:  "- First Album",
					})
				}

				for _, member := range artist.Members {
					if strings.Contains(strings.ToLower(member), query) {
						res = append(res, SearchSuggestion{
							ID:    artist.Id,
							Title: member,
							Type:  "- member",
						})
					}
				}

				if len(res) >= 15 {
					break
				}

			}
			if len(res) < 15 {
				for loc := range locations {
					if len(res) >= 15 {
						break
					}
					if strings.Contains(strings.ToLower(loc), query) {
						res = append(res, SearchSuggestion{
							Title: loc,
							Type:  "- location",
						})
					}
				}
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
			return
		} else if r.Method == http.MethodPost {
			query := r.FormValue("q")
			newArr := []methods.Artist{}
			mp := make(map[int]int)

			if query == "" {
				Display(w, artists, locations)
				return
			}

			query = strings.ToLower(query)
			for _, artist := range artists {
				if strings.Contains(strings.ToLower(artist.Name), query) {
					// continue to skip (phill collins has one member is himself)
					mp[artist.Id]++
					continue
				}

				creationDate := strconv.Itoa(artist.CreationDate)
				if creationDate == query {
					mp[artist.Id]++
				}

				if artist.FirstAlbum == query {
					mp[artist.Id]++
				}

				for _, member := range artist.Members {
					if strings.Contains(strings.ToLower(member), query) {
						mp[artist.Id]++
					}
				}

				for loc, id := range locations {
					if strings.Contains(strings.ToLower(loc), query) {
						for _, v := range id {
							if v == int8(artist.Id) {
								mp[artist.Id]++
							}
						}
					}
				}
			}

			for _, artist := range artists {
				if _, ok := mp[artist.Id]; ok {
					newArr = append(newArr, artist)
				}
			}

			Display(w, newArr, locations)
		} else {
			ErrorHandler(w, r, http.StatusMethodNotAllowed)
			return
		}
	}
}
