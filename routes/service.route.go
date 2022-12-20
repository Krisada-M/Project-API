package routes

import (
	"Restapi/controllers"

	"github.com/gin-gonic/gin"
)

// ServiceRoute is Route for non Authorization
func ServiceRoute(router *gin.RouterGroup) {
	service := router.Group("service")
	service.GET("/list", controllers.GetServiceList())
	service.POST("/add-booking", controllers.AddBooking())
}
