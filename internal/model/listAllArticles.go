package model

import (
	"log"
	"os"
	"strings"
)

// GenerateListOfAllArticles generates a list of all articles available in the "static/articles/" directory.
// It reads all files in the directory, filters out non-`.txt` files, and loads each article.
func GenerateListOfAllArticles() ([]Article, error) {
	// Define the directory where articles are stored.
	dir := "static/articles/"

	// Read the content of the specified directory.
	files, err := os.ReadDir(dir)
	if err != nil {
		// Log an error and terminate if the directory cannot be read.
		log.Fatalf("Error reading directory: %v", err)
	}

	// Create a slice to store the names of `.txt` files.
	var filenames []string

	// Loop through all the files in the directory.
	for _, file := range files {
		// Filter only `.txt` files and exclude directories.
		if !file.IsDir() && len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".txt" {
			// Add the file name (without the `.txt` extension) to the filenames slice.
			filenames = append(filenames, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}

	// Create a slice to hold the loaded articles.
	var articles []Article

	// Loop through the filenames to load each article.
	for _, filename := range filenames {
		// Load the content of each article by its filename (assuming LoadPage function handles it).
		article, _ := LoadPage(filename)

		// Append the loaded article to the articles slice.
		articles = append(articles, *article)
	}

	// Return the list of articles.
	return articles, nil
}
