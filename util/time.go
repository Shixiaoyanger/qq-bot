package util

import (
	"fmt"
	"time"
)

func GetTimeStampByMinute(timeStr string) int64 {
	timeStr = fmt.Sprintf("%s %s:00", time.Now().Format("2006-01-02"), timeStr)
	l, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, l)
	return t.Unix()
}
func GetTimeNow() int64 {
	return time.Now().Unix()
}
func GetExpireTime(timeStr string) int64 {
	return GetTimeStampByMinute(timeStr) - GetTimeNow()
}
