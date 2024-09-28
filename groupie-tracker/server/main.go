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
	}

	http.HandleFunc("/static/", handlers.StaticFileServer)
	http.Handle("/", handlers.HomeHandler(artists))
	http.Handle("/artists/", handlers.ArtistHandler(artists))

	fmt.Println("Server is running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}