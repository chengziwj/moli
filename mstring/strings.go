package mstring

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

//IsBlank 判断字符串是否为空
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

//Random 随机字符串
func Random(n int) string {
	b := make([]rune, n)
	l := len(defaultLetters)
	for i := range b {
		b[i] = defaultLetters[rand.Intn(l)]
	}
	return string(b)
}

//Reverse 字符串反转
func Reverse(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	runes := []rune(s)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//LeftPadInt 数字左填充
func LeftPadInt(s int64, l int) string {
	i := strconv.FormatInt(s, 10)
	ll := len(i)
	if ll >= l {
		return i
	}
	return strings.Repeat("0", l-ll) + i
}
