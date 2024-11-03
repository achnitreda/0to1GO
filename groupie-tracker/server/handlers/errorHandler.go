package handlers

import (
	"net/http"

	"web/methods"
)

// function that display the error page

func ErrorHandler(w http.ResponseWriter, _ *http.Request, status int) {
	w.WriteHeader(status)

	var title, message string
	switch status {
	case http.StatusNotFound:
		title = "Page Not Found"
		message = "The page you are looking for does not exist."
	case http.StatusInternalServerError:
		title = "Internal Server Error"
		message = "Something went wrong on our server."
	case http.StatusMethodNotAllowed:
		title = "Method Not Allowed"
		message = "The HTTP method used is not allowed for this endpoint."
	case http.StatusBadRequest:
		title = "Bad Request"
		message = "somthing wrong from your part"
	default:
		title = "Error"
		message = "An unexpected error occurred."
	}

	data := map[string]interface{}{
		"Title":  title,
		"Result": message,
		"Status": status,
	}

	methods.RenderTemplate(w, "error.html", data)
}
