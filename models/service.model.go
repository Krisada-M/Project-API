package models

import (
	"gorm.io/gorm"
)

// SalonService is Service in salon
type SalonService struct {
	gorm.Model
	Status    *string `json:"status" gorm:"default:'pending'"`
	Date      *string `json:"date"`
	TimeStart *string `json:"time_start"`
	TimeEnd   string  `json:"time_end"`
	Service   *string `json:"service"`
	BarberID  uint    `json:"barber_id"`
	UserID    uint    `json:"user_id"`
}

// ServiceList is service in salon
type ServiceList struct {
	ServiceName  *string       `json:"service_name" gorm:"primarykey"`
	ServiceSalon SalonService  `gorm:"ForeignKey:Service;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Service1     BarberProfile `gorm:"ForeignKey:Service1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Service2     BarberProfile `gorm:"ForeignKey:Service2;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Service3     BarberProfile `gorm:"ForeignKey:Service3;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Service4     BarberProfile `gorm:"ForeignKey:Service4;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// GetServiceList is service in salon
type GetServiceList struct {
	ServiceName *string `json:"service_name" `
}

// ServiceDetail is Service in salon
type ServiceDetail struct {
	Status    *string `json:"status"`
	Date      *string `json:"date"`
	TimeStart *string `json:"time_start"`
	TimeEnd   string  `json:"time_end"`
	Service   *string `json:"service"`
	BarberID  uint    `json:"barber_id"`
	UserID    uint    `json:"user_id"`
}

// GetServiceBooking with date and id barber
type GetServiceBooking struct {
	Date *string `json:"date"`
}

// ServiceMetaData metadata of service
type ServiceMetaData struct {
	ServiceID        uint `gorm:"primarykey"`
	LengthHair       string
	HairThickness    string
	UniquenessOfHair string
}

// AddBookingService model AddBooking
type AddBookingService struct {
	ServiceName      *string `json:"service_name"`
	Date             *string `json:"date"`
	TimeStart        *string `json:"time_start"`
	BarberID         uint    `json:"barber_id"`
	UserID           uint    `json:"user_id"`
	LengthHair       string  `json:"length_hair"`
	HairThickness    string  `json:"hair_thickness"`
	UniquenessOfHair string  `json:"uniqueness_of_hair"`
}
