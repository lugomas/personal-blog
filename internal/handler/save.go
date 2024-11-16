package handler

import (
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
	"time"
)

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
