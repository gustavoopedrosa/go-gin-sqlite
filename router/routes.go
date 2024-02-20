package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoopedrosa/go-gin-sqlite/handler"
)

const openingRoute = "opening"

// initializeRoutes initialize the routes of the API.
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET(openingRoute, handler.ShowOpeningHandler)

		v1.POST(openingRoute, handler.CreateOpeningHandler)

		v1.DELETE(openingRoute, handler.DeleteOpeningHandler)

		v1.PUT(openingRoute, handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
