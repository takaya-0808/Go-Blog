package handler

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/usecase"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("SECRETKEY")

type UserHandler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) Index(c *gin.Context) {

	_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})

	if err == nil {
		name := c.Param("name")
		user, err := uh.userUseCase.Search(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok", "user info": user})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not token"})
	}
}

func (uh userHandler) Show(c *gin.Context) {

	users, err := uh.userUseCase.Show()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "User's data": users})
}

func (uh userHandler) Register(c *gin.Context) {

	var user model.RegisterModel
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		return
	}
	token, err := uh.userUseCase.Add(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "miss register", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "token": token})
}

func (uh userHandler) Login(c *gin.Context) {

	var user model.LoginModel
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		return
	}
	token, err := uh.userUseCase.Check(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
