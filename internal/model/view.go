package model

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadPage(title string) (*Article, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile("static/articles/" + fileName)
	if err != nil {
		return nil, err
	}
	var article Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", title, err)
	}
	return &Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}
