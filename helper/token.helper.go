package helper

import (
	"Restapi/models"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	secretKey string = os.Getenv("SECRET_KEY")
)

// GenerateAllTokens is create JWT Token
func GenerateAllTokens(email string, firstName string, lastName string, userType string, userID string) (signedToken string, signedRefreshToken string, err error) {
	claims := &models.SignedDetail{
		Data: models.Payload{
			Email:     email,
			Firstname: firstName,
			Lastname:  lastName,
			UserType:  userType,
			UserID:    userID,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &models.SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken for Authorization
func ValidateToken(signToken string) (claims *models.SignedDetail, msg string) {
	token, err := jwt.ParseWithClaims(
		signToken,
		&models.SignedDetail{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*models.SignedDetail)

	if !ok {
		msg = fmt.Sprintf("token invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("Expire token")
		msg = err.Error()
		return
	}

	return claims, msg
}

// TokenGenerator for gennerate user_token
func TokenGenerator() string {
	token := make([]byte, 16)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}
