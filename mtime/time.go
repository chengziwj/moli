package mtime

import "time"

const DefaultLayout = "2006-01-02 15:04:05"

//Timestamp 获取unix时间戳
func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampInt() int {
	return int(Timestamp())
}

//Str2Unix 字符串转换为Unix时间戳
//默认格式 yyyy-MM-dd HH:mm:ss
func Str2Unix(str string) int64 {
	return ParseInt(DefaultLayout, str)
}

//Str2Unix 字符串转换为Unix时间戳
//默认格式 yyyy-MM-dd HH:mm:ss
func Str2UnixDefault(str string,defVal int64) int64 {
	return ParseIntDefault(DefaultLayout, str,defVal)
}

//ParseInt 字符串转换成时间戳
func ParseInt(layout, value string) int64 {
	return ParseIntLoc(layout, value, time.Local)
}

//ParseIntDefault 解析字符串为unix
func ParseIntDefault(layout, value string, defVal int64) int64 {
	t, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return defVal
	}
	return t.Unix()
}

//ParseIntLoc 字符串转换成时间戳
func ParseIntLoc(layout, value string, loc *time.Location) int64 {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//StartOfDay 获取当天开始时间戳
func StartOfDay(dt time.Time) int64 {
	year, month, day := dt.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, dt.Location()).Unix()
}

//EndOfDay 获取当天结束时间戳
func EndOfDay(dt time.Time) int64 {
	year, month, day := dt.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, dt.Location()).Unix() - 1
}
