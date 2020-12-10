package utils

import (
	"math/rand"
	"strings"
	"time"
)

func StringBuilder(p ...string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

var (
	letters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	letterLen = len(letters)
)

func RandNewStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(letterLen)
		data[i] = byte(letters[idx])
	}

	return string(data)
}
