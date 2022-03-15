package handler

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var secretKey := os.Getenv("SECRETKEY")

type BlogHandler interface {
	Show(c *gin.Context)
	TitleOneShow(c *gin.Context)
	TitleShow(c *gin.Context)
	GetArticle(c *gin.Context)
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

// GET function
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

func (bh blogHandler) GetArticle(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	artcle, err := bh.blogUsecase.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "artcle": artcle})
}

func (bu blogHandler) TitleOneShow(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	title, err := bu.blogUsecase.TitleOneShow(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "title": title})
}

// POST function
func (bh blogHandler) CreateArticle(c *gin.Context) {

	var blog model.CreateArticle
	if err := c.Bind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		return
	}
	errs := bh.blogUsecase.Create(&blog)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Bad Request", "error": "no article"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
