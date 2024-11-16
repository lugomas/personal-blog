package model

// Article represents a blog article with necessary metadata.
// It contains fields for the article's ID, title, content, and timestamps for creation and updates.
type Article struct {
	ID        string `json:"id"`         // Unique identifier for the article.
	Title     string `json:"title"`      // Title of the article.
	Content   []byte `json:"Content"`    // Content of the article stored as a byte slice.
	CreatedAt string `json:"created_at"` // Timestamp when the article was created.
	UpdatedAt string `json:"updated_at"` // Timestamp when the article was last updated.
}

var articlesPath = "static/articles/"
