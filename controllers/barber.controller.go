package controllers

import (
	"Restapi/config"
	"Restapi/helper"
	"Restapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BarberProfile by barber list
func BarberProfile(f string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			Count              int64
			barberID           = c.Param("bid")
			barber             = models.BarberProfile{}
			params             = models.GetServiceBooking{}
			barberAll          = []models.BarberProfile{}
			barberList         = []models.BarberProfileOnly{}
			store              = []models.SalonService{}
			resStore           = []models.ServiceDetail{}
			resBarberList      = []models.BarberProfileOnlyAndUser{}
			resBarberListAdmin = []models.BarberProfileOnlyAdmin{}
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

			c.JSON(http.StatusOK, helper.D{"barber_detail": resBarberList}.APIResponse())
			return
		case "B-ID":
			if err := c.BindJSON(&params); err != nil {
				c.JSON(http.StatusOK, helper.D{"Message": err}.APIResponse())
			}
			result := DB.Preload("Books").Table("barber_profiles").Where("id = ?", barberID).Find(&barber)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
				return
			}
			result = DB.Table("salon_services").Where("barber_id = ? AND date = ? AND status IN ?", barberID, params.Date, []string{"approve", "closed"}).Find(&store)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
				return
			}
			for _, v := range store {
				resStore = append(resStore, models.ServiceDetail{
					Status:    v.Status,
					Date:      v.Date,
					TimeStart: v.TimeStart,
					TimeEnd:   v.TimeEnd,
					Service:   v.Service,
					BarberID:  v.BarberID,
					UserID:    v.UserID,
				})
			}
			c.JSON(http.StatusOK, helper.D{"barber_detail": barber, "booking_detail": resStore}.APIResponse())
			return
		case "B-All":
			result := DB.Preload("Books").Table("barber_profiles").Order("id asc").Find(&barberAll)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
				return
			}

			for _, b := range barberAll {
				result = DB.Table("salon_services").Where("barber_id = ? AND date = ? AND status IN ?", b.ID, config.BookDate, []string{"approve", "closed"}).Find(&store).Count(&Count)
				if result.Error != nil {
					c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
					return
				}
				resBarberListAdmin = append(resBarberListAdmin, models.BarberProfileOnlyAdmin{
					ID:        b.ID,
					Name:      b.Name,
					Status:    b.Status,
					Gender:    b.Gender,
					Service1:  b.Service1,
					Service2:  b.Service2,
					Service3:  b.Service3,
					Service4:  b.Service4,
					BookInDay: strconv.FormatInt(Count, 10),
					AllBook:   strconv.Itoa(len(b.Books)),
				})
			}
			c.JSON(http.StatusOK, helper.D{"barber_detail": resBarberListAdmin}.APIResponse())
			return
		}
	}
}

// AddBarber admin
func AddBarber() gin.HandlerFunc {
	return func(c *gin.Context) {
		var barber = models.BarberProfile{}

		if err := c.BindJSON(&barber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}
		result := DB.Table("barber_profiles").Create(&barber)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		c.JSON(http.StatusOK, helper.D{"barber_id": barber.ID}.APIResponse())
		return
	}
}

// EditBarber admin
func EditBarber() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			barber   = models.BarberProfileOnlyAndUser{}
			barberID = c.Param("bid")
		)

		if err := c.BindJSON(&barber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}
		result := DB.Table("barber_profiles").Where("id = ?", barberID).Updates(models.BarberProfileOnlyAndUser{
			Name:     barber.Name,
			Gender:   barber.Gender,
			Service1: barber.Service1,
			Service2: barber.Service2,
			Service3: barber.Service3,
			Service4: barber.Service4,
		})

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		c.JSON(http.StatusOK, helper.D{"barber_id": barberID + " barber has update"}.APIResponse())
	}
}

// ChangeStatus for admin change status barber
func ChangeStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = models.ChangeStatus{}

		if err := c.BindJSON(&params); err != nil {
			c.JSON(http.StatusOK, helper.D{"Message": err}.APIResponse())
		}
		result := DB.Table("barber_profiles").Where("id = ?", params.BID).Update("status", params.Status)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
			return
		}
		c.JSON(http.StatusOK, helper.D{"barber": params.BID + " status has update"}.APIResponse())
	}
}

// DeleteBarber for admin remove barber
func DeleteBarber() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			barberID = c.Param("bid")
			barber   = models.BarberProfile{}
		)
		result := DB.Table("barber_profiles").Where("id = ?", barberID).Find(&barber)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "Barber not found"})
			return
		}
		result = DB.Delete(&barber)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "Barber not found"})
			return
		}

		c.JSON(http.StatusOK, helper.D{"barber": barberID + "delete success"}.APIResponse())
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
		result := DB.Table("barber_profiles").Where("name LIKE ? ", "%"+search.Keyword+"%").Find(&barberList)
		if search.Gender != "" && search.Service == "" {
			result = DB.Table("barber_profiles").Where("name LIKE ? AND gender = ?", "%"+search.Keyword+"%", search.Gender).Find(&barberList)
		}
		if search.Service != "" && search.Gender == "" {
			result = DB.Raw("SELECT * FROM barber_profiles WHERE name LIKE ? AND ? IN (service1,service2,service3,service4)", "%"+search.Keyword+"%", search.Service).Scan(&barberList)
		}
		if search.Gender != "" && search.Service != "" {
			result = DB.Raw("SELECT * FROM barber_profiles WHERE name LIKE ? AND gender = ? AND ? IN (service1,service2,service3,service4)", "%"+search.Keyword+"%", search.Gender, search.Service).Scan(&barberList)
		}
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
