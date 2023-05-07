package middlware

import (
	"golang_project_ecommerce/pkg/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("User_Athorization")
	if err != nil || tokenString == " " {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Error while fetching cookie",
		})
		return
	}

	claims, err := auth.Validatetoken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "error while validating",
		})
		return
	}
	//expiry time
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
		})
		return
	}

	c.Set("user_id", claims["sub"])

}
