package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadPage loads an article by its title from a file in the "static/articles/" directory.
// The article is deserialized from the stored JSON format into an Article struct.
func LoadPage(title string) (*Article, error) {
	// Define the file name based on the article's title with a `.txt` extension.
	fileName := title + ".txt"

	// Read the content of the article file.
	body, err := os.ReadFile("static/articles/" + fileName)
	if err != nil {
		// Return an error if the file cannot be read (e.g., if the file does not exist).
		return nil, err
	}

	// Declare a variable to hold the deserialized article.
	var article Article

	// Unmarshal the JSON content from the file into the article struct.
	err = json.Unmarshal(body, &article)
	if err != nil {
		// Return an error if the JSON parsing fails.
		return nil, fmt.Errorf("failed to parse %s: %w", title, err)
	}

	// Return the populated article struct.
	return &Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}
