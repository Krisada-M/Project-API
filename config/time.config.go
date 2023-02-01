package config

import (
	"strconv"
	"time"
)

// Monthlist is Thai Month
var Monthlist = [12]string{
	"มกราคม",
	"กุมภาพันธ์",
	"มีนาคม",
	"เมษายน",
	"พฤษภาคม",
	"มิถุนายน",
	"กรกฎาคม",
	"สิงหาคม",
	"กันยายน",
	"ตุลาคม",
	"พฤศจิกายน",
	"ธันวาคม",
}

// Day : current day
// Month : current month
// Year : current year
// Time : current time
var (
	tn       = time.Now()
	Day      = strconv.Itoa(tn.Day())
	Month    = Monthlist[tn.Month()-1]
	Year     = strconv.Itoa(tn.Year() + 543)
	Time     = tn.Format("15:04:05")
	Date     = (Day + " " + Month + " พ.ศ." + Year + " เวลา " + Time)
	BookDate = strconv.Itoa(tn.Year()) + "-" + strconv.Itoa(int(tn.Month())) + "-" + Day
)

// GetTime For Send e-mail
func GetTime(date string) (day, month, year string) {
	dateTime, _ := time.Parse("2006-01-02", date)
	Day := strconv.Itoa(dateTime.Day())
	Month := Monthlist[dateTime.Month()-1]
	Year := strconv.Itoa(dateTime.Year() + 543)
	return Day, Month, Year
}
