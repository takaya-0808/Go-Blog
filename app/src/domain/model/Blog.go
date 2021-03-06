package model

import "time"

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    string    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateArticle struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author"`
}

type TitlesShow struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
