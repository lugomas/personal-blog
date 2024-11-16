package handler

import (
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
)

// NewHandler handles requests to create a new page (article).
func NewHandler(w http.ResponseWriter, r *http.Request) {
	// Render the "new" template with an empty page (new article)
	page := model.Article{}
	RenderTemplate(w, "new", page)
}
