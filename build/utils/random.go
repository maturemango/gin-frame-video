package utils

import (
	"math/rand"
	"strings"
	"time"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func RandomVideoNo(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var b strings.Builder
	for i := 0; i < length; i ++ {
		b.WriteByte(chars[rand.Intn(len(chars))])
	}
	return b.String()
}