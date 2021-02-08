package moli

import "time"

func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampInt() int {
	return int(Timestamp())
}
