package models

import (
	"gorm.io/gorm"
)

// BarberProfile is Profile in salon
type BarberProfile struct {
	gorm.Model
	Name     *string        `json:"name,omitempty" validate:"required,min=2,max=100"`
	Gender   *string        `json:"gender,omitempty" validate:"required,min=2,max=100"`
	Status   *string        `json:"status" validate:"required,min=2,max=100"`
	Service1 *string        `json:"service1"`
	Service2 *string        `json:"service2"`
	Service3 *string        `json:"service3"`
	Service4 *string        `json:"service4"`
	Books    []SalonService `gorm:"ForeignKey:BarberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// BarberProfileOnly is Profile in salon
type BarberProfileOnly struct {
	gorm.Model
	Name     *string `json:"name,omitempty" validate:"required,min=2,max=100"`
	Gender   *string `json:"gender,omitempty" validate:"required,min=2,max=100"`
	Status   *string `json:"status" validate:"required,min=2,max=100"`
	Service1 *string `json:"service1"`
	Service2 *string `json:"service2"`
	Service3 *string `json:"service3"`
	Service4 *string `json:"service4"`
}

// BarberProfileOnlyAndUser is Profile in salon
type BarberProfileOnlyAndUser struct {
	ID       uint
	Name     *string `json:"name,omitempty" validate:"required,min=2,max=100"`
	Gender   *string `json:"gender,omitempty" validate:"required,min=2,max=100"`
	Status   *string `json:"status" validate:"required,min=2,max=100"`
	Service1 *string `json:"service1"`
	Service2 *string `json:"service2"`
	Service3 *string `json:"service3"`
	Service4 *string `json:"service4"`
}

// UserInBprofile in BarberProfileOnlyAndUser is Profile in salon
type UserInBprofile struct {
	Name     string
	Username *string
	Gender   *string
	UserType string
	Email    *string
}

// LiveSearch in barber list
type LiveSearch struct {
	Keyword string `json:"keyword"`
	Service string `json:"service"`
	Gender  string `json:"gender"`
}
