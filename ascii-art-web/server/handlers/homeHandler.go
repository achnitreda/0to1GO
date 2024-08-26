package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"web/methods"
)

type formret struct {
	Finput, Fbanner, Result string
}

func Homehandler(w http.ResponseWriter, r *http.Request) {

	const maxBodySize = 1024 * 3
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
	if err := r.ParseForm(); err != nil {
		if err.Error() == "http: request body too large" {
			ErrorHandler(w, r, http.StatusRequestEntityTooLarge)
		} else {
			ErrorHandler(w, r, http.StatusBadRequest)
		}
		return
	}

	template, err := template.ParseFiles("../client/templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	// handle requests depend on the method
	switch r.Method {
	// GET method
	case http.MethodGet:
		if r.URL.Path != "/" {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		} else {
			err = template.Execute(w, formret{Fbanner: "standard"})
			if err != nil {
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
		}
		// POST method
	case http.MethodPost:
		input := r.PostFormValue("input")
		font := r.PostFormValue("banner")
		if input == "" || !methods.IsValidBanner(font) {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		} else {
			banner := "banners/" + font + ".txt"
			validInput := methods.ValidInput(input)
			output := methods.ProccessOutput(validInput, methods.ProcessBanner(banner, validInput))
			// The use of buffer her is for not serve anything until the execution completed withount eny error
			var buf bytes.Buffer
			err = template.Execute(&buf, formret{Finput: input, Fbanner: font, Result: output})
			if err != nil {
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			} else {
				buf.WriteTo(w)
			}
		}
	default:
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}
