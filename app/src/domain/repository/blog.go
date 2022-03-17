package repository

import "Go-Blog/app/src/domain/model"

type BlogRepository interface {
	GetAllArticle() (*[]model.Article, error)
	GetOneArticle(id int) (*model.Article, error)
	TitleOneShow(id int) (*model.TitlesShow, error)
	TitleShow() (*[]model.TitlesShow, error)
	PostArticle(article *model.CreateArticle) error
}
