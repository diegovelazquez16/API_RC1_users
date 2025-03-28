package routes

import (
	"github.com/gin-gonic/gin"
	"holamundo/users/infraestructure/controllers"
)

func UserRoutes(
	router *gin.Engine,
	createController *controllers.UserCreateController,
	lastFiveController *controllers.UserLastFiveController,
	genderCountController *controllers.UserGenderCountController,
) {
	api := router.Group("/users")
	{
		api.POST("/", createController.Create)
		api.GET("/last-five", lastFiveController.GetLastFiveUsers)
		api.GET("/count", genderCountController.GetGenderCount)
	}
}
