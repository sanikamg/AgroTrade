package utils

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(data)

}

func StringToTime(date string) (time.Time, error) {
	layout := "2006-01-02"

	// Parse the string date using the specified layout
	returnDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	// Return the parsed time
	return returnDate, nil
}

func GenerateCouponCode(length int) string {
	// Define characters to be used in the coupon code
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Initialize random number generator with current time as seed
	rand.Seed(time.Now().UnixNano())

	// Generate a random coupon code of the specified length
	couponCode := strings.Builder{}
	for i := 0; i < length; i++ {
		couponCode.WriteByte(charSet[rand.Intn(len(charSet))])
	}

	return couponCode.String()
}
