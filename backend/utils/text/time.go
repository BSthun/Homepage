package text

import "time"

var BangkokTime *time.Location

func init() {
	BangkokTime, _ = time.LoadLocation("Asia/Bangkok")
}

func GetDateString() string {
	return time.Now().In(BangkokTime).Format("2006-01-02")
}
