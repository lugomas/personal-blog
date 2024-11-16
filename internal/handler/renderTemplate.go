package handler

import (
	"net/http"
	"text/template"
)

// RenderTemplate parses and executes a template with the provided page data.
func RenderTemplate(w http.ResponseWriter, templateName string, page interface{}) {
	// Parse the template file located at internal/tmpl/templates/
	tpl, err := template.ParseFiles("internal/tmpl/templates/" + templateName + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Execute the template and write the output to the response
	err = tpl.Execute(w, page)
	if err != nil {
		// If an error occurs while executing the template, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
