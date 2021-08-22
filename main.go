package main

import (
	"Go-Blog/app/config"
	"Go-Blog/app/src/infrastructure/database"
	"Go-Blog/app/src/interfaces/handler"
	"Go-Blog/app/src/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	userDatabases := database.NewUserDabase(config.Connection())
	userUseCases := usecase.NewUseUsecase(userDatabases)
	userHandler := handler.NewUserHandler(userUseCases)

	router := gin.Default()
	routerEngin := router.Group("/MyBlog")
	{
		v1 := routerEngin.Group("/api/v1")
		{
			v1.GET("/:name", userHandler.Index)
		}
	}

	router.Run(":8022")
}