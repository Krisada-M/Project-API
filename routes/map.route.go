package routes

import (
	"Restapi/controllers"
	"Restapi/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MapRoutes Map
func MapRoutes(router *gin.Engine) {
	app := router.Group("/app-json/v1/")
	AuthRoute(app)
	BarberRoute(app)
	ServiceRoute(app)
	app.Use(middleware.Authenticate())
	app.GET("checktype", controllers.CheckUserType())
	// user
	UserRoutes(app)
	//admin
	AdminRoutes(app)

	router.NoRoute(Stoproute)
}

// Stoproute is 423
func Stoproute(c *gin.Context) {
	c.JSON(http.StatusLocked, gin.H{"Stop": "You have found the quarantine station 🚩"})
	return
}
