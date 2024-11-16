package handler

import (
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
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
