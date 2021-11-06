package repository

import "Go-Blog/app/src/domain/model"

type BlogRepository interface {
	GetAllArticle() (*[]model.Article, error)
	GetOneArticle(id int) (*model.Article, error)
	PostArticle(*model.Article) error
}
