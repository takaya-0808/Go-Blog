package handler

import (
	"Go-Blog/app/src/usecase"

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

func (uh userHandler) Index(c *gin.Context) {}
