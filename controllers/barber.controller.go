package controllers

import (
	"Restapi/helper"
	"Restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BarberProfile by barber list
func BarberProfile(f string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			barberID      = c.Param("bid")
			barber        = models.BarberProfile{}
			barberList    = []models.BarberProfileOnly{}
			resBarberList = []models.BarberProfileOnlyAndUser{}
		)
		switch f {
		case "All":
			result := DB.Table("barber_profiles").Order("id asc").Find(&barberList)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
				return
			}
			for _, b := range barberList {
				resBarberList = append(resBarberList, models.BarberProfileOnlyAndUser{
					ID:       b.ID,
					Name:     b.Name,
					Status:   b.Status,
					Gender:   b.Gender,
					Service1: b.Service1,
					Service2: b.Service2,
					Service3: b.Service3,
					Service4: b.Service4,
				})
			}

			c.JSON(http.StatusOK, helper.D{"barber_detail": resBarberList, "test": barberList}.APIResponse())
			return
		case "B-ID":
			result := DB.Preload("Books").Table("barber_profiles").Where("id = ?", barberID).Find(&barber)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
				return
			}
			c.JSON(http.StatusOK, helper.D{"barber_detail": barber}.APIResponse())
			return
		}
	}
}

// AddBarber admin
func AddBarber() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			barber = models.BarberProfile{}
		)
		if err := c.BindJSON(&barber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}
		result := DB.Table("barber_profiles").Create(&barber)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		c.JSON(http.StatusOK, helper.D{"service_id": barber.ID}.APIResponse())
		return
	}
}

// LiveSearchBarber by barber list
func LiveSearchBarber() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			search        = models.LiveSearch{}
			barberList    = []models.BarberProfileOnly{}
			resBarberList = []models.BarberProfileOnlyAndUser{}
		)
		if err := c.BindJSON(&search); err != nil {
			BarberProfile("All")
			return
		}
		result := DB.Table("barber_profiles").Where("name LIKE ? AND gender LIKE ?", "%"+search.Keyword+"%", "%"+search.Gender+"%").Find(&barberList)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		for _, b := range barberList {
			resBarberList = append(resBarberList, models.BarberProfileOnlyAndUser{
				ID:       b.ID,
				Name:     b.Name,
				Status:   b.Status,
				Gender:   b.Gender,
				Service1: b.Service1,
				Service2: b.Service2,
				Service3: b.Service3,
				Service4: b.Service4,
			})
		}
		c.JSON(http.StatusOK, helper.D{"barber_detail": resBarberList}.APIResponse())
		return
	}
}
