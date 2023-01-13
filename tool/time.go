package tool

import "time"

func FormatTime() (formatTime string) {
	now := time.Now()
	formatTime = now.Format("2006-01-02")
	return formatTime
}
