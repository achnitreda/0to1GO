package methods

import (
	"net/http"
	"text/template"
	"time"
)
// Parsing all the templates one time
var templates = template.Must(template.ParseFiles(
	"../client/templates/pages/header.html",
	"../client/templates/pages/footer.html",
	"../client/templates/index.html",
	"../client/templates/artist.html",
	"../client/templates/error.html",
))

// function that execute templates

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	// Add common data
	if data == nil {
		data = make(map[string]interface{})
	}
	data["Year"] = time.Now().Year()

	// Execute the template
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
