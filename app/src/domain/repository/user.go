package repository

import 

type Usermodel interface {
	Get(int) (*model.UserModel, error)
	GetAll() []model.UserModel
	Delte(id int) error
	Add(user *model.UserModel) (id int, err error)
}