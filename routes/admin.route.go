package routes

import (
	"Restapi/controllers"
	"Restapi/middleware"

	"github.com/gin-gonic/gin"
)

// AdminRoutes is Route for user after login
func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/app-json/v1/admin", middleware.AccountAuthenticate("ADMIN"))
	admin.GET("/", func(c *gin.Context) {
		c.JSON(202, gin.H{"message": "Access granted for Admin"})
		return
	})
	admin.GET("/all-admin", controllers.GetAll("ADMIN"))
	admin.GET("/admin-id/:id", controllers.GetByID("ADMIN"))
	admin.GET("/all-user", controllers.GetAll("USER"))
	admin.GET("/user-id/:id", controllers.GetByID("USER"))
	// admin.GET("/get-by-barber", controllers.ManageService("BARBER_R"))
	// admin.GET("/get-all", controllers.ManageService("ADMIN_R"))
	// admin.POST("/add-service", controllers.ManageService("C"))
	// admin.PUT("/edit-service", controllers.ManageService("U"))
	// admin.DELETE("/delete-service/:sid", controllers.ManageService("D"))
}
