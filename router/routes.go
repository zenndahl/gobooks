
package router

import (
	"gobooks/handlers"

	//docs "gobooks/docs"

	"github.com/gin-gonic/gin"
	//swaggerfiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitHandler()

	basePath := "/api/v1"
//	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		//v1.GET("/books", handler.ListOpeningsHandler)
		//v1.GET("/books", handler.ShowOpeningHandler)
		v1.POST("/books", handlers.CreateBookHandler)
		//v1.DELETE("/books", handler.DeleteOpeningHandler)
		//v1.PUT("/books", handler.UpdateOpeningHandler)
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
