package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims)(string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err

	}
	return webtoken, nil
}
//Verifikasi Token JWT
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}