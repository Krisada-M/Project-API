package models

import jwt "github.com/dgrijalva/jwt-go"

// SignedDetail is model claims
type SignedDetail struct {
	Data Payload
	jwt.StandardClaims
}

// Payload Data
type Payload struct {
	Email     string
	Firstname string
	Lastname  string
	UserType  string
	UserID    string
}
