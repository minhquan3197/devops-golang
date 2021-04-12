package jwt

import (
	"project-golang/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Encrypt create jwt
func Encrypt(key string, payload interface{}) string {
	cfg := configs.Load()
	claims := jwt.MapClaims{}
	claims[key] = payload
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenKey := cfg.JWTKey
	token, _ := at.SignedString([]byte(tokenKey))
	return token
}

// Decrypt update jwt
func Decrypt(payload string) jwt.MapClaims {
	cfg := configs.Load()
	tokenKey := cfg.JWTKey
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(payload, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return nil
	}
	return claims
}
