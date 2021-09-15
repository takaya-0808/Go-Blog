package handler

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	name := c.Param("name")
	user, err := uh.userUseCase.Search(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "user info": user})
}

func (uh userHandler) Show(c *gin.Context) {

	users, err := uh.userUseCase.Show()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "user data": users})
}

func (uh userHandler) Register(c *gin.Context) {

	var user model.RegisterModel
	c.Bind(&user)
	token, err := uh.userUseCase.Add(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "miss register"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "token": token})
}

func (uh userHandler) Login(c *gin.Context) {

	var user model.LoginModel
	c.Bind(&user)
	token, err := uh.userUseCase.Check(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}
