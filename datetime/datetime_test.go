package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWithOption(t *testing.T) {
	dt := NewWithOption("2021-05-27 09:52:00",WithLoc(time.Local))
	fmt.Println(dt.Start(),dt.End())
}
