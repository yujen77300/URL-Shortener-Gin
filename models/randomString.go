package models

import (
	"strings"
	"time"
	"math/rand"
)

func RandomString(n int) string{
	var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	alphabetSize := len(alphabet)
	var sb strings.Builder

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}
	return sb.String()
}