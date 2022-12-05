package helper

import (
	"crypto/rand"
	"fmt"
	"log"
	mrand "math/rand"

	"golang.org/x/crypto/bcrypt"
)

// OTPGenerator is generate OTP.
func OTPGenerator() (string, string) {
	const (
		min = 16
		max = 20
	)
	var otp string
	randOtp, _ := rand.Prime(rand.Reader, mrand.Intn(max-min+1)+min)
	otp = splitByWidth(randOtp.String(), 6)

	if len(randOtp.String()) < 6 {
		otp = "0" + randOtp.String()
	}

	refNo := RefNoOTP(otp)
	return otp, refNo
}

// RefNoOTP is generate redno. OTP.
func RefNoOTP(OTP string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(OTP), 8)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// VerifyOTP is compare OTP.
func VerifyOTP(userOtp string, providedOtp string) (bool, string) {
	var (
		err   = bcrypt.CompareHashAndPassword([]byte(providedOtp), []byte(userOtp))
		check = true
		msg   = ""
	)

	if err != nil {
		msg = fmt.Sprintf("Otp is incorrect")
		check = false
	}

	return check, msg
}

func splitByWidth(str string, size int) string {
	var (
		strLength = len(str)
		splited   []string
		stop      int
	)
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		splited = append(splited, str[i:stop])
	}
	return splited[0]
}
