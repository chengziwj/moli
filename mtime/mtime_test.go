package mtime

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T)  {
	fmt.Println(ParseInt(DefaultLayout,"2021-03-23 12:19:24"))
}

func TestTimestamp(t *testing.T) {
	fmt.Println(TimestampInt())
}

func TestStartEnd(t *testing.T)  {
	dt := time.Now()
	fmt.Println(StartOfDay(dt),EndOfDay(dt))
	dt1,_ := time.ParseInLocation(DefaultLayout,"2021-04-30 17:39:22",time.Local)
	fmt.Println(StartOfDay(dt1),EndOfDay(dt1))
}
