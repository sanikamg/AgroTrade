package utils

import (
	"encoding/json"
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
