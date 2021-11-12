package handler

import (
	"github.com/gin-gonic/gin"
)

// var secretKey := os.Getenv("SECRETKEY")

type BlogHandler interface{}

type blogHandler struct{}

func (bh blogHandler) Show(c *gin.Context) {}
