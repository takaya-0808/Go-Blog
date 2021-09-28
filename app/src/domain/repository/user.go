package repository

import "Go-Blog/app/src/domain/model"

type UserRepository interface {
	Get(name string) (*model.UserModel, error)
	GetAll() []model.UserModel
	Add(name string, email string, pass string) error
	IDFindByName(name string) int
	IDFindByEmail(email string) int
	PassFindByName(name string) (string, error)
}

type BlogRepository interface {
	GetAllArticle() (*[]model.UserBlog, error)
	GetOneArticle(id int) (*model.UserBlog, error)
	PostArticle(*model.PostArticle) error
}
