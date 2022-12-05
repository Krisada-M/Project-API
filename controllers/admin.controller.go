package controllers

import (
	"Restapi/helper"
	"Restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAll for Admin
func GetAll(ut string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			personInfo = []models.User{}
			allPerson  = []models.ResponseUser{}
		)
		result := DB.Model(&models.User{}).Where("user_type = ?", ut).Find(&personInfo)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}
		for _, d := range personInfo {
			name := *d.Firstname + " " + *d.Lastname
			allPerson = append(allPerson, models.ResponseUser{
				ID:        d.ID,
				Name:      name,
				Username:  d.Username,
				UserType:  d.UserType,
				Email:     d.Email,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}
		count := len(allPerson)
		c.JSON(http.StatusOK, helper.D{"_count": count, "all_person": allPerson}.APIResponse())
		return
	}
}

// GetByID for find by ID
func GetByID(ut string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userID = c.Param("id")
			user   models.User
			person = models.ResponseUser{}
		)
		switch ut {
		case "ADMIN":
			if err := helper.MatchAdminTypeToID(c, userID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": "ADMIN not found"})
				return
			}
			break
		case "USER":
			if err := helper.MatchUserTypeToID(c, userID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": "User not found"})
				return
			}
			break
		}
		result := DB.Model(&models.User{}).Where("id = ? AND user_type = ?", userID, ut).Find(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}
		name := *user.Firstname + " " + *user.Lastname
		person = models.ResponseUser{
			ID:        user.ID,
			Name:      name,
			Username:  user.Username,
			UserType:  user.UserType,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		c.JSON(http.StatusOK, helper.D{"person_detail": person}.APIResponse())
		return
	}
}
