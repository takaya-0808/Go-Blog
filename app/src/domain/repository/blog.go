package repository

import "Go-Blog/app/src/domain/model"

type BlogRepository interface {
	GetAllArticle() (*[]model.Article, error)
	GetOneArticle(id int) (*model.Article, error)
	TitleShow() (*[]model.TitlesShow, error)
	PostArticle(*model.CreateArticle) error
}
