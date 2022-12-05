package config

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Mailer : models go-mail
var Mailer *gomail.Dialer

// ConnectMailer : function connect e-mail
func ConnectMailer(username, password string) {
	// Envload()
	port, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))
	mailer := gomail.NewDialer(
		os.Getenv("MAILER_HOST"),
		port,
		username,
		password,
	)
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	Mailer = mailer
}
