package main

import (
	"fmt"
	"log"
	"net/http"
	"web/handlers"
	"web/methods"
)

func main() {
	artistURL := "https://groupietrackers.herokuapp.com/api/artists"

	var artists []methods.Artist
	err := methods.FetchParser(artistURL, &artists)
	if err != nil {
		log.Print(err)
		return
	}

	locations, err := methods.GetLocations()
	if err != nil {
		log.Print(err)
		return
	}

	http.Handle("/", handlers.HomeHandler(artists, locations))
	http.Handle("/search/", handlers.SearchHandler(artists, locations))
	http.Handle("POST /filter/", handlers.Filter(artists, locations))
	http.Handle("/artists/", handlers.ArtistHandler(artists))
	http.HandleFunc("/static/", handlers.StaticFileServer)

	fmt.Println("Server is running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
