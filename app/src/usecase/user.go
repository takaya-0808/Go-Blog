package usecase

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

type UserUseCase interface {
	Search(name string) (*model.UserModel, error)
	Check(user model.LoginModel) (string, error)
	Show() ([]model.UserModel, error)
	Add(user model.RegisterModel) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
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

func (uu userUseCase) Check(user model.LoginModel) (string, error) {

	hashpass, err := uu.userRepository.PassFindByName(user.UserName)
	if err != nil {
		return "0", err
	}

	err = uu.CheckHashPass(hashpass, user.UserPassWord)
	if err != nil {
		return "0", err
	}

	AccessToken := jwt.New(jwt.GetSigningMethod("HS256"))
	AccessToken.Claims = jwt.MapClaims{
		"user": "admin",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenString, err := AccessToken.SignedString([]byte(secretKey))
	return tokenString, nil
}

func (uu userUseCase) Show() ([]model.UserModel, error) {

	var users []model.UserModel
	return users, nil
}

func (uu userUseCase) Add(user model.RegisterModel) (string, error) {

	passHash := uu.CreateHashPass(user.UserPassWord)
	if uu.userRepository.IDFindByName(user.UserName) != 0 || uu.userRepository.IDFindByEmail(user.UserEmail) != 0 {
		return "0", errors.New("used name or email")
	}

	err := uu.userRepository.Add(user.UserName, user.UserEmail, passHash)
	if err != nil {
		return "0", err
	}

	AccessToken := jwt.New(jwt.GetSigningMethod("HS256"))
	AccessToken.Claims = jwt.MapClaims{
		"user": "admin",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenString, err := AccessToken.SignedString([]byte(secretKey))
	return tokenString, nil
}

func (uu userUseCase) CheckHashPass(hashpass string, pass string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}

func (uu userUseCase) CreateHashPass(pass string) string {

	passhash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "0"
	}
	passHash := string(passhash)
	return passHash
}

func (uu userUseCase) VerifyToken(tokenString string) (*jwt.Token, error) {

	var token *jwt.Token
	var err error
	return token, err
}
