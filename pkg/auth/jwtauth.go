package auth

import (
	"errors"
	"fmt"
	"golang_project_ecommerce/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int) (map[string]string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetJWTConfig()))

	if err != nil {
		return nil, errors.New("JWT token generating is failed")
	}
	return map[string]string{"accessToken": tokenString}, nil
}

func Validatetoken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.GetJWTConfig()), nil
	})
	if err != nil || !token.Valid {
		return jwt.MapClaims{}, errors.New("not valid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

		return jwt.MapClaims{}, errors.New("can't parse claims")

	}

	return claims, nil
}

// sub as phone number
func GenerateJWTPhn(phn string) (map[string]string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": phn,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetJWTConfig()))

	if err != nil {
		return nil, errors.New("JWT token generating is failed")
	}
	return map[string]string{"accessToken": tokenString}, nil
}
