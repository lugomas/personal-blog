package handler

import (
	"fmt"
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
	"text/template"
)

// HomeHandler handles requests to display the home page, which lists all articles.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Generate the list of all articles
	articles, err := model.GenerateListOfAllArticles()
	if err != nil {
		// If there is an error generating the list of articles, return an internal server error
		http.Error(w, fmt.Sprintf("Error generating list of articles: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the home page template
	tpl, err := template.ParseFiles("internal/tmpl/templates/home.html")
	if err != nil {
		// If there is an error loading the template, return an internal server error
		http.Error(w, fmt.Sprintf("Error loading template: %v", err), http.StatusInternalServerError)
		return
	}

	// Execute the template with the list of articles and write the output to the response
	err = tpl.Execute(w, articles)
	if err != nil {
		// If there is an error executing the template, return an internal server error
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
