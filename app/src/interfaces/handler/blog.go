package handler

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var secretKey := os.Getenv("SECRETKEY")

type BlogHandler interface {
	Show(c *gin.Context)
	TitleShow(c *gin.Context)
	CreateArticle(c *gin.Context)
}

type blogHandler struct {
	blogUsecase usecase.BlogUseCase
}

func NewBlogHandler(bu usecase.BlogUseCase) BlogHandler {
	return &blogHandler{
		blogUsecase: bu,
	}
}

func (bh blogHandler) Show(c *gin.Context) {

	artcles, err := bh.blogUsecase.Index()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "blog": artcles})
}

func (bh blogHandler) TitleShow(c *gin.Context) {

	title, err := bh.blogUsecase.TitleShow()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"titles": title})
}

func (bu blogHandler) CreateArticle(c *gin.Context) {

	var blog model.CreateArticle
	if err := c.Bind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
