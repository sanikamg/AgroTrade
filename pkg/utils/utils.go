package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(data)

}
