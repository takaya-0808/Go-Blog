package usecase

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
	"errors"
)

type UserUseCase interface {
	Search(name string) (*model.UserModel, error)
	Add(user model.RegisterModel) (string, error)
	Show() ([]model.UserModel, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUseUsecase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) Search(name string) (*model.UserModel, error) {

	user, err := uu.userRepository.Get(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Add(user model.RegisterModel) (string, error) {

	token := "0"
	if user.UserName == "hoge" && user.UserPassWord == "hoge" {
		token = "1"
		return token, nil
	}
	return token, errors.New("error")
}

func (uu userUseCase) Show() ([]model.UserModel, error) {

	var users []model.UserModel
	return users, nil
}
