package routes

import (
	"api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		decks := main.Group("decks")
		{
			decks.GET("/", controllers.ShowDecks)
		}
	}

	return router
}
