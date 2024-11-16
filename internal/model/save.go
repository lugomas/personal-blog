package model

import (
	"encoding/json"
	"fmt"
	"os"
)

func (p *Article) Save() error {
	fileName := p.Title + ".txt"
	newData, err := json.MarshalIndent(Article{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, "", "  ")
	if err != nil {
		return fmt.Errorf("could not save article: %v", err)
	}
	return os.WriteFile("static/articles/"+fileName, newData, 0600)
}
