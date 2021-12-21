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
func LeftPadInt(i int64, l int) string {
	s := strconv.FormatInt(i, 10)
	ll := len(s)
	if ll >= l {
		return s
	}
	return strings.Repeat("0", l-ll) + s
}

func Cut(str string, l int, tail string) string {
	runes := []rune(str)
	tails := []rune(tail)
	ll := len(runes)
	lt := len(tails)

	if ll <= l {
		return string(runes)
	}

	//字符串超过限定长度, lc + lt = l
	runes = runes[:l-lt]
	runes = append(runes, tails...)
	return string(runes)
}

//WrapString 分割字符串，并且前后添加指定字符
func WrapString(str, sep, s string) string {
	return Wrap(strings.Split(str, sep), s)
}

//Wrap 将元素前后添加指定字符
func Wrap(arr []string, s string) string {
	var rs []string
	for _, v := range arr {
		rs = append(rs, s+v+s)
	}
	return strings.Join(rs, ",")
}

//JoinInt int数组转换成字符串
func JoinInt(nums []int, sep string) string {
	b := make([]string, len(nums))
	for i, v := range nums {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

//JoinUint int数组转换成字符串
func JoinUint(nums []uint64, sep string) string {
	b := make([]string, len(nums))
	for i, v := range nums {
		b[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(b, sep)
}

//JoinInt64 int64数组转换成字符串
func JoinInt64(nums []int64, sep string) string {
	b := make([]string, len(nums))
	for i, v := range nums {
		b[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(b, sep)
}
