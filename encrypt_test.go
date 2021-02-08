package moli

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	s := "12345"
	fmt.Println(Md5(s))
}
