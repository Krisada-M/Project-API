package routes

import (
	"Restapi/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MapRoutes Map
func MapRoutes(router *gin.Engine) {

	AuthRoute(router)
	BarberRoute(router)
	ServiceRoute(router)
	router.Use(middleware.Authenticate())
	// user
	UserRoutes(router)
	//admin
	AdminRoutes(router)

	router.NoRoute(Stoproute)
}

// Stoproute is 423
func Stoproute(c *gin.Context) {
	c.JSON(http.StatusLocked, gin.H{"Stop": "You have found the quarantine station 🚩"})
	return
}
