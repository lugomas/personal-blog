package model

import (
	"log"
	"os"
	"strings"
)

func GenerateListOfAllArticles() ([]Article, error) {
	dir := "static/articles/"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	var filenames []string
	for _, file := range files {
		// Filter only .txt files
		if !file.IsDir() && len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".txt" {
			filenames = append(filenames, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}

	var articles []Article
	for _, filename := range filenames {
		article, _ := LoadPage(filename)
		articles = append(articles, *article)
	}
	return articles, nil
}
