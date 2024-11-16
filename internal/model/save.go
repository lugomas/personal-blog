package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Save saves an article to a file with the title as the file name. The article is serialized into JSON format
// and saved to the "static/articles/" directory.
func (p *Article) Save() error {
	// Define the file name based on the article's title with a `.txt` extension.
	fileName := p.Title + ".txt"

	// Serialize the article into a formatted JSON structure with indentation.
	newData, err := json.MarshalIndent(Article{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, "", "  ")

	// If there is an error during the marshalling process, return the error.
	if err != nil {
		return fmt.Errorf("could not save article: %v", err)
	}

	// Write the serialized JSON data to a file in the "static/articles/" directory.
	// Set the file permissions to `0600`, which means read and write permissions for the owner only.
	return os.WriteFile("static/articles/"+fileName, newData, 0600)
}
