package routes

import (
	"Restapi/controllers"
	"Restapi/middleware"

	"github.com/gin-gonic/gin"
)

// AdminRoutes is Route for user after login
func AdminRoutes(router *gin.RouterGroup) {
	admin := router.Group("admin", middleware.AccountAuthenticate("ADMIN"))
	admin.GET("/", func(c *gin.Context) {
		c.JSON(202, gin.H{"message": "Access granted for Admin"})
		return
	})
	admin.GET("/all-admin", controllers.GetAll("ADMIN"))
	admin.GET("/admin-id/:id", controllers.GetByID("ADMIN"))
	admin.GET("/all-user", controllers.GetAll("USER"))
	admin.GET("/user-id/:id", controllers.GetByID("USER"))
	admin.GET("/barber-detail", controllers.BarberProfile("B-All"))
	admin.GET("/booking-pending", controllers.GetBookingByStatus("pending"))
	admin.GET("/booking-approve", controllers.GetBookingByStatus("approve"))
	admin.GET("/booking-closed", controllers.GetBookingByStatus("closed"))
	admin.GET("/booking-unapproved", controllers.GetBookingByStatus("unapproved"))
	admin.POST("/barber-status", controllers.ChangeStatus())
	admin.POST("/barber-delete/:bid", controllers.DeleteBarber())
	admin.POST("/booking-status", controllers.UpdateStatusBooking())
	admin.POST("/booking-remove/:sid", controllers.DeleteBooking())
}
