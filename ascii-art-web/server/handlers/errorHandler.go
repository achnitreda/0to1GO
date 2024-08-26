package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title, Result string
	Status        int
}

func ErrorHandler(w http.ResponseWriter, _ *http.Request, status int) {
	tmpl, err := template.ParseFiles("../client/templates/error.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	var title, message string
	switch status {
	case http.StatusNotFound:
		title = "Page Not Found"
		message = "The page you are looking for does not exist."
	case http.StatusInternalServerError:
		title = "Internal Server Error"
		message = "Something went wrong on our server."
	case http.StatusBadRequest:
		title = "Bad Request"
		message = "The request could not be understood by the server."
	case http.StatusMethodNotAllowed:
		title = "Method Not Allowed"
		message = "The HTTP method used is not allowed for this endpoint."
	case http.StatusRequestEntityTooLarge:
		title = "Request body too large"
		message = "The request body is too large."
	default:
		title = "Error"
		message = "An unexpected error occurred."
	}
	err = tmpl.Execute(w, PageData{Title: title, Result: message, Status: status})
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
