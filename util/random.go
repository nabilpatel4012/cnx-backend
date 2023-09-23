package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Random String generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomUser() string {
	return RandomString(10)
}

func RandomPhone() string {
	return randomdata.PhoneNumber()
}

func RandomEmail() string {
	return randomdata.Email()
}

func RandomAddress() string {
	return randomdata.Address()
}

func RandomOrder() int {
	return randomdata.Number(20)
}
