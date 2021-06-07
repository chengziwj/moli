package mdate

import "time"

//StartOfDay 获取当天开始时间戳
func StartOfDay(dt time.Time) int64 {
	year, month, day := dt.Date()
	return ZeroPoint(year, month, day, dt.Location()).Unix()
}

//EndOfDay 获取当天结束时间戳
func EndOfDay(dt time.Time) int64 {
	year, month, day := dt.Date()
	return ZeroPoint(year, month, day+1, dt.Location()).Unix() - 1
}

//StartOfMonth 返回月份第一秒时间戳
func StartOfMonth(dt time.Time) int64 {
	year, month, _ := dt.Date()
	return ZeroPoint(year, month, 1, dt.Location()).Unix()
}

//EndOfMonth 返回月份最后一秒时间戳
func EndOfMonth(dt time.Time) int64 {
	year, month, _ := dt.Date()
	return ZeroPoint(year, month+1, 1, dt.Location()).Unix() - 1
}

//ZeroPoint 返回零点时间
func ZeroPoint(year int, month time.Month, day int, loc *time.Location) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}
