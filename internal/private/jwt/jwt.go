package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Encrypt create jwt
func Encrypt(key string, payload interface{}) string {
	claims := jwt.MapClaims{}
	claims[key] = payload
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenKey := os.Getenv("JWT_KEY")
	token, _ := at.SignedString([]byte(tokenKey))
	return token
}

// Decrypt update jwt
func Decrypt(payload string) jwt.MapClaims {
	tokenKey := os.Getenv("JWT_KEY")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(payload, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return nil
	}
	return claims
}
