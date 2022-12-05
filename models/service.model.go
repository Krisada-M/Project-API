package models

import (
	"gorm.io/gorm"
)

// SalonService is Service in salon
type SalonService struct {
	gorm.Model
	Status    *string `json:"status"`
	Date      *string `json:"date"`
	TimeStart *string `json:"time_start"`
	TimeEnd   *string `json:"time_end"`
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

// ServiceDetail is Service in salon
type ServiceDetail struct {
	Status    *string `json:"status"`
	Date      *string `json:"date"`
	TimeStart *string `json:"time_start"`
	TimeEnd   *string `json:"time_end"`
	Service   uint    `json:"service"`
	UserID    uint    `json:"user_id"`
}
