package mtime

import "time"

const DefaultLayout = "2006-01-02 15:04:05"

func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampInt() int {
	return int(Timestamp())
}

func Parse(layout, value string) int64 {
	return ParseLocation(layout, value, time.Local)
}

func ParseLocation(layout, value string, loc *time.Location) int64 {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}
