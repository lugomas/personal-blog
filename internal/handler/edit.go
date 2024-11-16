package handler

import (
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
)

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
