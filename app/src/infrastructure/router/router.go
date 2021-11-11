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
	userDatabases := database.NewUserDabase(config.Connection())
	userUseCases := usecase.NewUseUsecase(userDatabases)
	userHandler := handler.NewUserHandler(userUseCases)

	routerEngin := router.Group("/MyBlog")
	{
		v1 := routerEngin.Group("/api/v1")
		{
			v1.GET("/", userHandler.Show)
			v1.GET("/:name", userHandler.Index)
		}
		auth := routerEngin.Group("/api/auth")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/register", userHandler.Register)
		}

	}
	return router
}
