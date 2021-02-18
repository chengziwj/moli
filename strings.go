package moli

import (
	"math/rand"
	"strings"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	l := len(defaultLetters)
	for i := range b {
		b[i] = defaultLetters[rand.Intn(l)]
	}
	return string(b)
}
