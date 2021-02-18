package moli

import (
	"fmt"
	"testing"
)

func TestIsBlank(t *testing.T) {
	s := "  "
	if !IsBlank(s) {
		t.Error(s)
	}
}

func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(0))
	fmt.Println(RandomString(10))
}