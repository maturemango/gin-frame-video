package utils

import (
	"math/rand"
	"strings"
	_"time"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func RandomNo(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var b strings.Builder
	for i := 0; i < length; i ++ {
		b.WriteByte(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func RandomCode(length int) string {
	chars := "0123456789"
	var b strings.Builder
	for i := 0; i < length; i ++ {
		b.WriteByte(chars[rand.Intn(len(chars))])
	}
	return b.String()
}