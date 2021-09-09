package usecase

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Search(name string) (*model.UserModel, error)
	Check(user model.RegisterModel) (string, error)
	Show() ([]model.UserModel, error)
	Add(user model.RegisterModel) (string, error)
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

func (uu userUseCase) Check(user model.RegisterModel) (string, error) {

	token := "1"
	err := uu.userRepository.Check(user.UserName, user.UserEmail)

	if err != nil {
		return "0", err
	}

	if user.UserName == "hoge" || user.UserEmail == "hoge" {
		token = "0"
		return token, errors.New("used name or email")
	}
	return token, nil
}

func (uu userUseCase) Show() ([]model.UserModel, error) {

	var users []model.UserModel
	return users, nil
}

func (uu userUseCase) Add(user model.RegisterModel) (string, error) {

	token := "1"
	passhash, err := bcrypt.GenerateFromPassword([]byte(user.UserPassWord), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	passHash := string(passhash)
	err = uu.userRepository.Add(user.UserName, user.UserEmail, passHash)
	if err != nil {
		return "0", err
	}
	return token, nil
}
