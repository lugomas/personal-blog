package model

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   []byte `json:"Content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
