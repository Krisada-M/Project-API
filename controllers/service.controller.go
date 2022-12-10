package controllers

import (
	"Restapi/helper"
	"Restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetServiceList is get all service in salon
func GetServiceList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service = []models.GetServiceList{}
		result := DB.Table("service_lists").Find(&service)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, helper.D{"service_list": service}.APIResponse())
		return
	}
}

// AddBooking is user add booking service
func AddBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			service = models.AddBookingService{}
		)
		if err := c.BindJSON(&service); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}
		salonService := models.SalonService{
			Date:      service.Date,
			TimeStart: service.TimeStart,
			Service:   service.ServiceName,
			BarberID:  service.BarberID,
			UserID:    service.UserID,
		}
		result := DB.Table("salon_services").Create(&salonService)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		result = DB.Table("service_meta_data").Create(&models.ServiceMetaData{
			ServiceID:        salonService.ID,
			LengthHair:       service.LengthHair,
			HairThickness:    service.HairThickness,
			UniquenessOfHair: service.UniquenessOfHair,
		})
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		c.JSON(http.StatusOK, helper.D{"booking_id": salonService.ID}.APIResponse())
		return
	}
}
