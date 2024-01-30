
package router

import (
	"gobooks/handlers"

	docs "gobooks/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/books", handlers.ListBooksHandler)
		v1.GET("/books", handlers.ShowBookHandler)
		v1.POST("/books", handlers.CreateBookHandler)
		v1.DELETE("/books", handlers.DeleteBookHandler)
		v1.PUT("/books", handlers.UpdateBookHandler)
	}

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
