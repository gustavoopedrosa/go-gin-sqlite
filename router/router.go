package router

import "github.com/gin-gonic/gin"

// Initialize initializes and run the router.
func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run()
}
