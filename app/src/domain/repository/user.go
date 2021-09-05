package repository

import "Go-Blog/app/src/domain/model"

type UserRepository interface {
	Get(name string) (*model.UserModel, error)
	GetAll() []model.UserModel
	// Delte(id int) error
	// Add(user *model.UserModel) (id int, err error)
	RegisterCheck(name string, email string) error
}

type BlogRepository interface{}
