package controllers

import (
	"Restapi/database"
	"Restapi/helper"
	"Restapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
)

// DB is database connection
var (
	DB                = database.DB
	validate          = validator.New()
	googleOauthConfig *oauth2.Config
	oauthStateString  = "pseudo-random"
)

// CheckEmail is check email user
func CheckEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userEmail models.EmailUser
			count     int64
		)
		if err := c.BindJSON(&userEmail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Register : " + err.Error()})
			return
		}
		DB.Table("users").Where("email = ?", userEmail.Email).Count(&count)

		c.JSON(http.StatusOK, helper.D{"count": count}.APIResponse())
	}
}

// Register for those who apply for an account
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user  models.User
			check int64
		)

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Register : " + err.Error()})
			return
		}

		validattionErr := validate.Struct(user)
		if validattionErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Message": "Register : " + validattionErr.Error()})
			return
		}

		emailCheck := DB.Table("users").Where(models.User{Email: user.Email}).Count(&check)
		if check > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Email already", "Gorm Message": emailCheck.Error})
			return
		}

		userCheck := DB.Table("users").Where(models.User{Username: user.Username}).Count(&check)
		if check > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "User already", "Gorm Message": userCheck.Error})
			return
		}

		passScore, passMsg := helper.ValidatePassword(*user.Password, nil)
		if passScore < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "password is " + passMsg})
			return
		}

		password, _ := helper.HashPassword(*user.Password)

		user.Password = &password

		result := DB.Table("users").Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "", "Gorm Message": result.Error})
			return
		}

		c.JSON(http.StatusCreated, helper.D{"ID": user.ID}.APIResponse())
	}
}

// Login for user login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user              = models.UserLogin{}
			foundUser         = new(models.User)
			loginType *string = nil
		)

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Login : " + err.Error()})
			return
		}

		validattionErr := validate.Struct(user)
		if validattionErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Message": "Login : " + validattionErr.Error()})
			return
		}

		switch user.LoginType {
		case "username":
			loginType = user.Username
			break
		case "email":
			loginType = user.Email

			if emailPass := helper.ValidateEmailAddress(*user.Email); emailPass {
				c.JSON(http.StatusBadRequest, gin.H{"Message": "Please fill your Email"})
				return
			}
			break
		default:
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Login type is Incorrect"})
			return
		}

		result := DB.Where(user.LoginType+" = ?", loginType).Find(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "is Incorrect" + result.Error.Error()})
			return
		}

		if foundUser.UserType == "" {
			c.JSON(http.StatusNotFound, gin.H{"Message": "user not found"})
			return
		}

		passwordIsValid, msg := helper.VerifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusBadRequest, gin.H{"Message": msg})
			return
		}

		id := strconv.FormatUint(uint64(foundUser.ID), 10)
		token, _, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Firstname, *foundUser.Lastname, foundUser.UserType, id)

		c.JSON(http.StatusOK, helper.D{"accessToken": token}.APIResponse())
	}
}

// VerifyAccount for check account
func VerifyAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			verify    models.VerifyUser
			foundUser models.User
			count     int64
		)

		if err := c.BindJSON(&verify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Verify : " + err.Error()})
			return
		}

		validattionErr := validate.Struct(verify)
		if validattionErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Message": "Verify : " + validattionErr.Error()})
			return
		}

		if emailPass := helper.ValidateEmailAddress(*verify.Email); emailPass {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Verify : Please fill your Email"})
			return
		}

		DB.Table("users").Where("email = ?", verify.Email).Find(&foundUser).Count(&count)
		if count == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "User Not Found"})
			return
		}
		otp, refNo := helper.OTPGenerator()
		send := helper.Mailer{}
		send.OtpSendMail(foundUser, otp, refNo[len(refNo)-5:])

		c.JSON(http.StatusOK, helper.D{"user_id": foundUser.ID, "ref_no": refNo}.APIResponse())
	}
}

// VerifyOTP for check OTP
func VerifyOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var otpcheck models.OTP

		if err := c.BindJSON(&otpcheck); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "OTP_error : " + err.Error()})
			return
		}

		validattionErr := validate.Struct(otpcheck)
		if validattionErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Message": "OTP_error : " + validattionErr.Error()})
			return
		}

		passwordIsValid, msg := helper.VerifyOTP(*otpcheck.OTP, *otpcheck.RefNo)
		if passwordIsValid != true {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "OTP_error : " + msg})
			return
		}

		if otpcheck.UserID == "" {
			c.JSON(http.StatusNotFound, gin.H{"Message": "No user_id"})
			return
		}

		userToken := helper.TokenGenerator()

		result := DB.Model(&models.User{}).Where("id = ?", otpcheck.UserID).Update("otp_token", userToken)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error})
			return
		}

		c.JSON(http.StatusOK, helper.D{"user_token": userToken}.APIResponse())
	}
}

// SetPassword for users who want to change their password
func SetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userToken = c.Request.Header.Get("userToken")
			userPass  models.UserPassword
			user      models.User
		)

		if userToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "No userToken"})
			return
		}

		if err := c.BindJSON(&userPass); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "SetPassword : " + err.Error()})
			return
		}

		validattionErr := validate.Struct(userPass)
		if validattionErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Message": "SetPassword : " + validattionErr.Error()})
			return
		}

		passScore, passMsg := helper.ValidatePassword(*userPass.Password, nil)
		if passScore < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "password is " + passMsg})
			return
		}

		password, _ := helper.HashPassword(*userPass.Password)
		result := DB.Model(&models.User{}).Where("otp_token = ?", userToken).Find(&user)
		if user.ID == 0 || result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "user not found"})
			return
		}

		result = DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{Password: &password, OTPToken: "No Token"})
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, helper.D{}.APIResponse())
	}
}

// ChangePassword for users who forgot their password.
func ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userPass  models.UserPassword
			foundUser models.User
			userID    = c.GetString("userID")
		)

		if err := c.BindJSON(&userPass); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "ChangePassword" + err.Error()})
			return
		}

		validattionErr := validate.Struct(userPass)
		if validattionErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Message": "ChangePassword" + validattionErr.Error()})
			return
		}

		passScore, passMsg := helper.ValidatePassword(*userPass.NewPassword, nil)
		if passScore < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"Message": "password is " + passMsg})
			return
		}

		result := DB.Model(&models.User{}).Where("id = ?", userID).Find(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		passwordIsValid, msg := helper.VerifyPassword(*userPass.OldPassword, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusBadRequest, gin.H{"Message": msg})
			return
		}

		password, _ := helper.HashPassword(*userPass.NewPassword)

		foundUser.Password = &password

		result = DB.Save(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, helper.D{}.APIResponse())
	}
}

// GetUser is present User
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			foundUser models.User
			userID    = c.GetString("userID")
			user      = models.ResponseUser{}
		)

		result := DB.Preload("Books").Where("id = ?", userID).Find(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
			return
		}

		name := *foundUser.Firstname + " " + *foundUser.Lastname
		user = models.ResponseUser{
			ID:        foundUser.ID,
			Name:      name,
			Username:  foundUser.Username,
			Gender:    foundUser.Gender,
			UserType:  foundUser.UserType,
			Email:     foundUser.Email,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
			Books:     foundUser.Books,
		}
		c.JSON(http.StatusOK, helper.D{"user_detail": user}.APIResponse())
	}
}

// UpdateUser is present User
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userID   = c.GetString("userID")
			editUser = models.UpdateUser{}
		)

		if err := c.BindJSON(&editUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		result := DB.Table("users").Where("id = ?", userID).Updates(models.User{Firstname: editUser.Firstname, Lastname: editUser.Lastname, Email: editUser.Email, Gender: editUser.Gender})
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, helper.D{"user_detail": "Update success"}.APIResponse())
	}
}

// DeleteUser is present User
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userID    = c.GetString("userID")
			send      = helper.Mailer{}
			foundUser = models.User{}
		)

		result := DB.Where("id = ?", userID).Find(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
			return
		}

		send.DeleteUserSendMail(foundUser)
		result = DB.Delete(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
			return
		}

		c.JSON(http.StatusOK, helper.D{"user_detail": userID + "delete success"}.APIResponse())
	}
}
