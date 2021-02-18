package moli

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	s := "12345"
	fmt.Println(Md5(s))
}

func TestPassword(t *testing.T)  {
	pwd := "q1w2e3r4t5"
	h,_ := PasswordHash(pwd)
	fmt.Println(h)
	fmt.Println(PasswordVerify(h,pwd))
}