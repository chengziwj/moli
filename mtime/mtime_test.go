package mtime

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T)  {
	fmt.Println(Parse(DefaultLayout,"2021-03-23 12:19:24"))
}

func TestTimestamp(t *testing.T) {
	fmt.Println(TimestampInt())
}
