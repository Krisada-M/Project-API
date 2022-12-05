package routes

import (
	"Restapi/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoute is Route for non Authorization
func AuthRoute(router *gin.Engine) {
	general := router.Group("/app-json/v1/")
	general.GET("/", func(c *gin.Context) {
		c.JSON(202, gin.H{"success": "Access granted for General User"})
		return
	})
	general.POST("/check-email", controllers.CheckEmail())
	general.POST("/register", controllers.Register())
	general.POST("/login", controllers.Login())
	general.POST("/verify-account", controllers.VerifyAccount())
	general.POST("/verify-otp", controllers.VerifyOTP())
	general.POST("/set-password", controllers.SetPassword())
}
