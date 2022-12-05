package helper

import (
	"errors"
	"fmt"

	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/nbutton23/zxcvbn-go"
	"golang.org/x/crypto/bcrypt"
)

// CheckUserType is check access permissions
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType != role {
		err = errors.New(userType + "Unauthorized to access this resource" + role)
		return err
	}
	return err
}

// MatchUserTypeToID is check userType and userID
func MatchUserTypeToID(c *gin.Context, id string) (err error) {
	userType := c.GetString("userType")
	uid := c.GetString("userID")
	err = nil

	if userType == "USER" && uid != id {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

// MatchAdminTypeToID is check userType and userID
func MatchAdminTypeToID(c *gin.Context, id string) (err error) {
	userType := c.GetString("userType")
	uid := c.GetString("userID")
	err = nil

	if userType == "ADMIN" && uid != id {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

// ValidateEmailAddress is validate E-mail format
func ValidateEmailAddress(address string) bool {
	_, err := mail.ParseAddress(address)

	if err != nil {
		return true
	}

	return false
}

// HashPassword generate in bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword is compare password.
func VerifyPassword(userPassword, providedPassword string) (bool, string) {
	var (
		check = true
		msg   = ""
	)

	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		msg = fmt.Sprintf("password not match")
		check = false
	}

	return check, msg
}

// ValidatePassword for check PasswordStrength
func ValidatePassword(password string, userInput []string) (int, string) {
	var msg string

	passwordErr := zxcvbn.PasswordStrength(password, userInput)
	switch passwordErr.Score {
	case 0:
		msg = "too guessable"
		break
	case 1:
		msg = "very guessable"
		break
	case 2:
		msg = "somewhat guessable"
		break
	case 3:
		msg = "safely unguessable"
		break
	case 4:
		msg = "very unguessable"
		break
	}

	return passwordErr.Score, msg
}
