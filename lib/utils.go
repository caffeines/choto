package lib

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqABCDEFGHIJKLrstuvwxyzMNOPQRSTUVWXYZ1234567890")

// RandStringRunes will return n length string
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
