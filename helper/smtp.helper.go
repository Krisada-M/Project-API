package helper

import (
	"log"

	"Restapi/config"
	"Restapi/models"

	"gopkg.in/gomail.v2"
)

// Mailer models
type Mailer struct{}

// Send is send E-mail
func (m *Mailer) Send(message *gomail.Message) {
	message.SetHeader("From", "noreply@hairapppointment.com")

	if err := config.Mailer.DialAndSend(message); err != nil {
		log.Panicln("[Mailer] ", err)
	}
}

// OtpSendMail is send E-mail with Otp
func (m *Mailer) OtpSendMail(foundUser models.User, otp, ref string) {
	message := gomail.NewMessage()
	message.SetHeader("From", "noreply@hairapppointment.com")
	message.SetHeader("To", *foundUser.Email)
	message.SetHeader("Subject", "Hello! "+*foundUser.Firstname)
	message.SetBody("text/html", "<h1 style='margin:0; color:blue;'> Hairapppointment </h1>"+
		"<br> <h2>สวัสดี "+*foundUser.Firstname+" "+*foundUser.Lastname+"</h2>"+
		"<br> <center><h3>รหัส OTP ของคุณคือ "+otp+"<br>รหัสยืนยันของคุณคือ "+ref+"</h3></center>")

	m.Send(message)
}

// DeleteUserSendMail is send E-mail with Otp
func (m *Mailer) DeleteUserSendMail(foundUser models.User) {
	message := gomail.NewMessage()
	message.SetHeader("From", "noreply@hairapppointment.com")
	message.SetHeader("To", *foundUser.Email)
	message.SetHeader("Subject", "Hello! "+*foundUser.Firstname)
	message.SetBody("text/html", "<h1 style='margin:0; color:blue;'> Hairapppointment </h1>"+
		"<br> <h2>สวัสดี "+*foundUser.Firstname+" "+*foundUser.Lastname+"</h2>"+
		"<br> <center><h3>ระบบได้ทำการยกเลิกการลงทะเบียนกับทางร้านเรียบร้อย</h3></center>")

	m.Send(message)
}
