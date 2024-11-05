package handlers

import (
	"encoding/json"
	"fmt"
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
		query := r.URL.Query().Get("q")

		fmt.Println(r.URL)

		var res []SearchSuggestion
		var person methods.Artist

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
			if creationDate == query {
				res = append(res, SearchSuggestion{
					ID:    artist.Id,
					Title: artist.Name,
					Type:  "- artist/band",
				})
			}

			if artist.FirstAlbum == query {
				res = append(res, SearchSuggestion{
					ID:    artist.Id,
					Title: artist.Name,
					Type:  "- artist/band",
				})
			}

			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), query) {
					res = append(res, SearchSuggestion{
						ID:    artist.Id,
						Title: artist.Name,
						Type:  "- member",
					})
				}
			}
		}

		for loc, id := range locations {
			if strings.Contains(strings.ToLower(loc), query) {
				// fmt.Println(loc, id)
				if len(id) == 1 {
					res = append(res, SearchSuggestion{
						ID:    int(id[0]),
						Title: loc,
						Type:  "-" + strconv.Itoa(int(id[0])),
					})
				} else {
					for _, v := range id {
						url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%v", v)
						_ = methods.FetchParser(url, &person)
						res = append(res, SearchSuggestion{
							ID:    person.Id,
							Title: person.Name,
							Type:  "- artist/band",
						})
					}
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
