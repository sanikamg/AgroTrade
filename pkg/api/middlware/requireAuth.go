package middlware

import (
	"errors"
	"fmt"
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

func AutheticatePhn(c *gin.Context) {
	RequireAuth(c, "Phone_Authorization")
}

func RequireAuth(c *gin.Context, authname string) {
	fmt.Println(authname)
	tokenString, err := c.Cookie(authname)
	fmt.Println(tokenString)

	if err != nil || tokenString == " " {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Error while fetching cookie",
			"error":      err,
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

// to fetch id from jwt
func GetId(c *gin.Context, authname string) (float64, error) {
	cookie, err := c.Request.Cookie(authname)
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

// to fetch phone number from jwt
func GetPhn(c *gin.Context, authname string) (string, error) {
	cookie, err := c.Request.Cookie(authname)
	if err != nil {
		return " ", errors.New("can't find cookie")
	}

	tokenString := cookie.Value
	claims, err := auth.Validatetoken(tokenString)
	if err != nil {
		return " ", errors.New("can't validate cookie")
	}

	phn, ok := claims["sub"].(string)
	if !ok {
		return " ", errors.New("can't find phn")
	}
	return phn, nil
}
