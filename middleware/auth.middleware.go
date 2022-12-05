package middleware

import (
	"Restapi/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authenticate is JWT authorization
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}
		clientToken = strings.Split(clientToken, "Bearer ")[1]
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Data.Email)
		c.Set("firstname", claims.Data.Firstname)
		c.Set("lastname", claims.Data.Lastname)
		c.Set("userID", claims.Data.UserID)
		c.Set("userType", claims.Data.UserType)
		c.Next()
	}
}

// AccountAuthenticate for User authorization
func AccountAuthenticate(ut string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		err := helper.CheckUserType(c, ut)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := helper.MatchUserTypeToID(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
}
