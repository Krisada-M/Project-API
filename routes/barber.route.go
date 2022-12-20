package routes

import (
	"Restapi/controllers"

	"github.com/gin-gonic/gin"
)

// BarberRoute is Route for non Authorization
func BarberRoute(router *gin.RouterGroup) {
	barber := router.Group("barber")
	barber.POST("/search", controllers.LiveSearchBarber())
	barber.GET("/list", controllers.BarberProfile("All"))
	barber.POST("/:bid", controllers.BarberProfile("B-ID"))
	barber.POST("/add", controllers.AddBarber())
}
