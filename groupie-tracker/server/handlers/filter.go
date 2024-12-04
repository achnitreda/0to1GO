package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"web/methods"
)

func Filter(artists []methods.Artist, locations map[string][]int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newArr := []methods.Artist{}

		location := r.FormValue("location")
		minCreationDate := r.FormValue("minCd")
		maxCreationDate := r.FormValue("maxCd")
		nums := r.Form["members"]
		minAlbumDate := r.FormValue("minFa")
		maxAlbumDate := r.FormValue("maxFa")

		for _, n := range nums {
			n, err := strconv.Atoi(n)
			if err != nil {
				ErrorHandler(w, r, http.StatusBadRequest)
				return
			}

			for _, artist := range artists {
				if len(artist.Members) == n {
					newArr = append(newArr, artist)
				}
			}
		}

		if len(newArr) == 0 {
			newArr = append(newArr, artists...)
		}

		if minCreationDate != "" && maxCreationDate != "" {
			min, err := strconv.Atoi(minCreationDate)
			if err != nil {
				ErrorHandler(w, r, http.StatusBadRequest)
				return
			}

			max, err := strconv.Atoi(maxCreationDate)
			if err != nil {
				ErrorHandler(w, r, http.StatusBadRequest)
				return
			}

			for i := 0; i < len(newArr); i++ {
				if newArr[i].CreationDate < min || newArr[i].CreationDate > max {
					newArr = append(newArr[0:i], newArr[i+1:]...)
					i--
				}
			}
		}

		if location != "" {
			for i := 0; i < len(newArr); i++ {
				if _, ok := locations[location]; !ok {
					ErrorHandler(w, r, http.StatusBadRequest)
					return
				}

				if !isExists(int8(newArr[i].Id), locations[location]) {
					newArr = append(newArr[0:i], newArr[i+1:]...)
					i--
				}
			}
		}

		if minAlbumDate != "" && maxAlbumDate != "" {
			min, err := strconv.Atoi(minAlbumDate)
			if err != nil {
				ErrorHandler(w, r, http.StatusBadRequest)
				return
			}

			max, err := strconv.Atoi(maxAlbumDate)
			if err != nil {
				ErrorHandler(w, r, http.StatusBadRequest)
				return
			}

			for i := 0; i < len(newArr); i++ {
				firstAlbum, _ := strconv.Atoi(strings.Split(newArr[i].FirstAlbum, "-")[2])
				if firstAlbum < min || firstAlbum > max {
					newArr = append(newArr[0:i], newArr[i+1:]...)
					i--
				}
			}
		}

		Display(w, newArr, locations)
	}
}

func isExists(el int8, arr []int8) bool {
	for _, ele := range arr {
		if ele == el {
			return true
		}
	}
	return false
}
