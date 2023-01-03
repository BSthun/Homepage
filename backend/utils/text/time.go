package text

import "time"

var bangkokTime *time.Location

func init() {
	bangkokTime, _ = time.LoadLocation("Asia/Bangkok")
}

func GetDateString() string {
	return time.Now().In(bangkokTime).Format("2006-01-02")
}
