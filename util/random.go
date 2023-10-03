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
	return randomdata.Number(1, 3)
}
func RandomPrice() int {
	return randomdata.Number(100, 1000)
}

func RandomDeliveryTime() time.Time {

	// Set a time range (adjust as needed)
	minTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	maxTime := time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)

	// Generate a random duration within the time range
	timeRange := maxTime.Sub(minTime)
	randomDuration := time.Duration(rand.Int63n(int64(timeRange)))

	// Add the random duration to the minimum time to get a random delivery time
	randomDeliveryTime := minTime.Add(randomDuration)

	return randomDeliveryTime

}

func RandomOrderStatus() string {
	return randomdata.StringSample("Started", "Pending", "Accepted", "Work In Progress", "Done", "Delivered", "Cancelled")
}
