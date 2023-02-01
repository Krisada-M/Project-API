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

// ApproveBookingSendMail is send E-mail with Otp
func (m *Mailer) ApproveBookingSendMail(foundUser models.User, date, service, timeStart string) {
	Day, Month, Year := config.GetTime(date)
	message := gomail.NewMessage()
	message.SetHeader("From", "noreply@hairapppointment.com")
	message.SetHeader("To", *foundUser.Email)
	message.SetHeader("Subject", "Hello! "+*foundUser.Firstname)
	message.SetBody("text/html", "<h1 style='margin:0; color:green;'> Hairapppointment </h1>"+
		"<br> <h2>สวัสดี "+*foundUser.Firstname+" "+*foundUser.Lastname+"</h2>"+
		"<br> <center><h2 style='margin:0; color:green;'>บริการ "+service+" ที่คุณจองได้รับการยืนยันแล้ว</h2><br>"+
		"<h3>โปรดมาในวันที่ "+Day+" "+Month+" "+Year+" เวลา "+timeStart+" น.</h3><br>"+
		"<span>** โปรดมาก่อนเวลานัดหมายประมาณ 15 นาที **<span>"+"</center>")

	m.Send(message)
}

// UnApproveBookingSendMail is send E-mail with Otp
func (m *Mailer) UnApproveBookingSendMail(foundUser models.User, date, service, timeStart string) {
	Day, Month, Year := config.GetTime(date)
	message := gomail.NewMessage()
	message.SetHeader("From", "noreply@hairapppointment.com")
	message.SetHeader("To", *foundUser.Email)
	message.SetHeader("Subject", "Hello! "+*foundUser.Firstname)
	message.SetBody("text/html", "<h1 style='margin:0; color:red;'> Hairapppointment </h1>"+
		"<br> <h2>สวัสดี "+*foundUser.Firstname+" "+*foundUser.Lastname+"</h2>"+
		"<br> <center><h2 style='margin:0; color:red;'>บริการ "+service+" ที่คุณจองได้รับการปฏิเสธจากแอดมิน</h2><br>"+
		"<h3>เนื่องจากใหนวันที่ "+Day+" "+Month+" "+Year+" เวลา "+timeStart+" น. มีผู้ใช้บริการท่านอื่นจองมาก่อนหน้านี้แล้ว</h3><br>"+
		"<span>ขออภัยในความไม่สะดวกโปรดทำการจองใหม่อีกครั้งผ่านลิ้งค์ด้านล่าง<span><br>"+
		"<span>https://dream-salon.vercel.app/<span>"+"</center>")

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
