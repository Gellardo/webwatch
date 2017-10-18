package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandomString(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(buf)
}
