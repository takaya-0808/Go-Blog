package handler

import (
	"Go-Blog/app/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Index(c *gin.Context)
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