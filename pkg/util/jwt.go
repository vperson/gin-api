package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte("gin-api")

// Claims Custom claims structure
type Claims struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token
func GenerateToken(userId uint, username, password string, expire time.Duration) (string, error) {
	expirationTime := time.Now().Add(expire)
	claims := &Claims{
		UserId:   userId,
		Username: username,
		Password: password,
	}
	if expire != 0 {
		claims.StandardClaims = jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses a JWT token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
