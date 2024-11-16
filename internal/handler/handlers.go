package handler

import (
	"fmt"
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
	"text/template"
	"time"
)

// ViewHandler handles requests to view a page by its title.
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the page title from the URL path (removing "/view/")
	title := r.URL.Path[len("/view/"):]

	// Load the page content using the title
	page, err := model.LoadPage(title)
	if err != nil {
		// If the page does not exist, redirect to the edit page
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	// Render the "view" template with the loaded page data
	RenderTemplate(w, "view", page)
}

// EditHandler handles requests to edit a page by its title.
func EditHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the page title from the URL path (removing "/edit/")
	title := r.URL.Path[len("/edit/"):]

	// Attempt to load the page
	page, err := model.LoadPage(title)
	if err != nil {
		// If the page doesn't exist, create a new empty page structure
		page = &model.Article{
			ID:        "",
			Title:     title,
			Content:   nil,
			CreatedAt: "",
			UpdatedAt: "",
		}
	}

	// Render the "edit" template with the page data (empty if the page is new)
	RenderTemplate(w, "edit", page)
}

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

// SaveHandler handles saving a page, either creating a new page or updating an existing one.
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the page title from the URL path (removing "/save/")
	title := r.URL.Path[len("/save/"):]

	// If no title is passed in the URL, get it from the form data
	if title == "" {
		title = r.FormValue("title")
		body := r.FormValue("body")
		articleID := time.Now().Format("200601021504105") // Generate a unique ID using the current timestamp
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		updatedAt := time.Now().Format("2006-01-02 15:04:05")

		// Create a new article with the provided title, body, and timestamps
		article := &model.Article{
			ID:        articleID,
			Title:     title,
			Content:   []byte(body),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		// Save the new article to the data store
		err := article.Save()
		if err != nil {
			// If saving the article fails, return an internal server error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect the user to the view page for the newly created article
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
		return
	}

	// If the title is provided, update the existing article
	body := r.FormValue("body")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	// Load the existing article by its title
	existingArticle, err := model.LoadPage(title)
	if err != nil {
		// If loading the article fails, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new article with the updated content and timestamp
	article := &model.Article{
		ID:        existingArticle.ID,
		Title:     title,
		Content:   []byte(body),
		CreatedAt: existingArticle.CreatedAt,
		UpdatedAt: updatedAt,
	}

	// Save the updated article
	err = article.Save()
	if err != nil {
		// If saving the updated article fails, return an internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user to the view page for the updated article
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	return
}

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

// AdminHandler handles requests to the admin page, which lists all articles.
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// Generate the list of all articles
	articles, err := model.GenerateListOfAllArticles()
	if err != nil {
		// If there is an error generating the list of articles, return an internal server error
		http.Error(w, fmt.Sprintf("Error generating list of articles: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the admin page template
	tpl, err := template.ParseFiles("internal/tmpl/templates/admin.html")
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

// NewHandler handles requests to create a new page (article).
func NewHandler(w http.ResponseWriter, r *http.Request) {
	// Render the "new" template with an empty page (new article)
	page := model.Article{}
	RenderTemplate(w, "new", page)
}

// LogoutHandler handles logout requests and prompts the browser to re-authenticate.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Set the WWW-Authenticate header to prompt for credentials again
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization required"`)
	// Return a 401 Unauthorized status to trigger the browser's authentication prompt
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
