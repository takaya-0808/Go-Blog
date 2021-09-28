package model

type PostArticle struct {
	id      int    `json:"Id"`
	title   string `json:"Title"`
	context string `json:"Context"`
}
