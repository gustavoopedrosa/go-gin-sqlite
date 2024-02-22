package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/gustavoopedrosa/go-gin-sqlite/docs"
	"github.com/gustavoopedrosa/go-gin-sqlite/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const openingRoute = "opening"

// initializeRoutes initialize the routes of the API.
func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET(openingRoute, handler.ShowOpeningHandler)

		v1.POST(openingRoute, handler.CreateOpeningHandler)

		v1.DELETE(openingRoute, handler.DeleteOpeningHandler)

		v1.PUT(openingRoute, handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
