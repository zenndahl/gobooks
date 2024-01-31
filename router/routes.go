package router

import (
	//docs "gobooks/docs"

	"gobooks/handlers"

	"github.com/gin-gonic/gin"
	//swaggerfiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitHandler()

	basePath := "/api/v1"
	//docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/books", handlers.ListBooksHandler)
		v1.GET("/book", handlers.ShowBookHandler)
		v1.GET("/random-book", handlers.ShowRandomBookHandler)
		v1.POST("/book", handlers.CreateBookHandler)
		v1.DELETE("/book", handlers.DeleteBookHandler)
		v1.PUT("/book", handlers.UpdateBookHandler)
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
