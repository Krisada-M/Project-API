package controllers

import (
	"Restapi/helper"
	"Restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// CheckUserType is filter type for client
func CheckUserType() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"type": c.GetString("userType")})
		return
	}
}

// GetBookingByStatus for admin manage
func GetBookingByStatus(status string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookingList = []models.SalonService{}

		result := DB.Table("salon_services").Where("status = ?", status).Order("created_at DESC").Find(&bookingList)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, helper.D{"booking_list": bookingList}.APIResponse())
		return
	}
}

// UpdateStatusBooking is update approve booking
func UpdateStatusBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			service = models.UpadateBooking{}
			result  *gorm.DB
		)
		if err := c.BindJSON(&service); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if service.TimeEnd != "" {
			result = DB.Table("salon_services").Where("id = ?", service.ServiceID).Select("status", "time_end").Updates(models.SalonService{Status: service.Status, TimeEnd: service.TimeEnd})
		} else {
			result = DB.Table("salon_services").Where("id = ?", service.ServiceID).Update("status", service.Status)
		}

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, helper.D{"booking_id": service.ServiceID + " has update"}.APIResponse())
		return
	}
}

// DeleteBooking is delate booking for admin
func DeleteBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ServiceID = c.Param("sid")
			count     int64
			service   = models.SalonService{}
		)
		result := DB.Table("salon_services").Where("id = ?", ServiceID).Find(&service).Count(&count)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		if count < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Booking not found"})
			return
		}

		DB.Delete(&service)

		c.JSON(http.StatusOK, helper.D{"booking_id": ServiceID + " has delete"}.APIResponse())
		return
	}
}
