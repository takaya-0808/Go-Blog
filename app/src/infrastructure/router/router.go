package router

import (
	"Go-Blog/app/config"
	"Go-Blog/app/src/infrastructure/database"
	"Go-Blog/app/src/interfaces/handler"
	"Go-Blog/app/src/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouting() *gin.Engine {

	router := gin.Default()

	// User data
	userDatabases := database.NewUserDabase(config.Connection())
	userUseCases := usecase.NewUseUsecase(userDatabases)
	userHandler := handler.NewUserHandler(userUseCases)

	// Blog data
	blogDatabase := database.NewBlogDabase(config.Connection())
	blogUseCases := usecase.NewBlogUseCase(blogDatabase)
	blogHandler := handler.NewBlogHandler(blogUseCases)

	routerEngin := router.Group("/MyBlog")
	{
		v1 := routerEngin.Group("/api/v1")
		{
			v1.GET("", userHandler.Show)
			v1.GET("/:name", userHandler.Index)
		}
		auth := routerEngin.Group("/api/user")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/register", userHandler.Register)
		}
		article := routerEngin.Group("/api/article")
		{
			article.GET("", blogHandler.Show)
			article.POST("", blogHandler.CreateArticle)
			article.GET("/:id", blogHandler.GetArticle)
			article.GET("/titles", blogHandler.TitleShow)
		}

	}
	return router
}
