package models

import (
	"time"

	"gorm.io/gorm"
)

// User general models
type User struct {
	gorm.Model
	Firstname *string        `json:"firstname,omitempty" validate:"required,min=2,max=100"`
	Lastname  *string        `json:"lastname,omitempty" validate:"required,min=2,max=100"`
	Gender    *string        `json:"gender"`
	Username  *string        `json:"username,omitempty" gorm:"unique;" validate:"required,min=2,max=100"`
	Email     *string        `json:"email,omitempty" gorm:"unique;" validate:"email,required"`
	Password  *string        `json:"password" validate:"required,min=8"`
	UserType  string         `json:"user_type" gorm:"default:'USER'"`
	OTPToken  string         `json:"otp_token" gorm:"column:otp_token;size:32;default:'No Token'"`
	Books     []SalonService `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// UpdateUser general models
type UpdateUser struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Gender    *string `json:"gender"`
	Email     *string `json:"email"`
}

// EmailUser general model emailuser
type EmailUser struct {
	Email *string `json:"email"`
}

// ResponseUser general models
type ResponseUser struct {
	ID        uint
	Name      string
	Username  *string
	Gender    *string
	UserType  string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []SalonService
}

// UserLogin for Login function
type UserLogin struct {
	LoginType string  `json:"login_type" validate:"required,eq=email|eq=username"`
	Email     *string `json:"email,omitempty" validate:"required_without=Username"`
	Username  *string `json:"username,omitempty" validate:"required_without=Email"`
	Password  *string `json:"password" validate:"required,min=8"`
}

// VerifyUser for VerifyAccount function
type VerifyUser struct {
	UserID    string  `json:"user_id"`
	Email     *string `json:"email,omitempty" validate:"email,required"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
}

// OTP for Verufy OTP function
type OTP struct {
	UserID string  `json:"user_id" validate:"required"`
	RefNo  *string `json:"ref_no" validate:"required,min=5"`
	OTP    *string `json:"otp" validate:"required,min=6"`
}

// UserPassword for Change or Forgot password
type UserPassword struct {
	Password    *string `json:"password"`
	OldPassword *string `json:"old_password"`
	NewPassword *string `json:"new_password"`
}

// UserOTP for Change or Forgot password
type UserOTP struct {
	TokenID string `json:"token_id" validate:"required"`
	Token   string `json:"token" validate:"required"`
}
