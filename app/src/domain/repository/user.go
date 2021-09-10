package repository

import "Go-Blog/app/src/domain/model"

type UserRepository interface {
	Get(name string) (*model.UserModel, error)
	GetAll() []model.UserModel
	FindByName(name string) (string, error)
	Add(name string, email string, pass string) error
}

type BlogRepository interface{}
