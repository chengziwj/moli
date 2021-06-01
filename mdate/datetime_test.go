package mdate

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWithOption(t *testing.T) {
	dt := New("2021-05-27 09:52:00", WithLoc(time.Local))
	fmt.Println(dt.Start(), dt.End())
	fmt.Println(dt.FormatDay(), dt.FormatMonth(), dt.Digit(), dt.DayDigit())
}

func TestMonth(t *testing.T) {
	fmt.Println(New("2021-06", WithLayout("2006-01")).ToString())
}

func TestMonthSec(t *testing.T) {
	dt := Now()
	fmt.Println(dt.LastSecOfMonth(), dt.FirstSecOfMonth())
}
