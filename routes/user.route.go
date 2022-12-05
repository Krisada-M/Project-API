package routes

import (
	"Restapi/controllers"
	"Restapi/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoutes is Route for user after login
func UserRoutes(router *gin.Engine) {
	user := router.Group("/app-json/v1/user", middleware.AccountAuthenticate("USER"))
	user.GET("/", controllers.GetUser())
	user.POST("/update", controllers.UpdateUser())
	// user.POST("/service", controllers.UserService("R"))
	// user.POST("/u-service/:sid", controllers.UserService("U"))
	// user.POST("/d-service/:sid", controllers.UserService("D"))
	user.POST("/change-password", controllers.ChangePassword())
	user.POST("/cancel-subscription", controllers.DeleteUser())
}
