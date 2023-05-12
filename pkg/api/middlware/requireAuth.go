package middlware

import (
	"errors"
	"golang_project_ecommerce/pkg/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser(c *gin.Context) {
	RequireAuth(c, "User_Authorization")
}

func AuthenticateAdmin(c *gin.Context) {
	RequireAuth(c, "Admin_Authorization")
}

func RequireAuth(c *gin.Context, authname string) {

	tokenString, err := c.Cookie(authname)
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

	c.Set("userId", claims["sub"])

}

func GetId(c *gin.Context, authname string) (float64, error) {
	cookie, err := c.Request.Cookie("Admin_Authorization")
	if err != nil {
		return 0, errors.New("can't find cookie")
	}

	tokenString := cookie.Value
	claims, err := auth.Validatetoken(tokenString)
	if err != nil {
		return 0, errors.New("can't validate cookie")
	}

	id, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("can't find id")
	}

	return id, nil
}
