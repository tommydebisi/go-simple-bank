package utils

import (
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// generates a random word of length n
func RandomString(n int) string {
	var sb strings.Builder

	letterLength := len(alphabet)

	for i := 0; i < n; i++ {
		// get and store a random char between 0 - len-1
		char := alphabet[rand.Intn(letterLength)]

		// write single char to buffer
		sb.WriteByte(char)
	}

	return sb.String()
}

// generates a random owner
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currency := []string{"USD", "CAD", "NGN", "EUR"}
	
	currLength := len(currency)

	return currency[rand.Intn(currLength)]
}