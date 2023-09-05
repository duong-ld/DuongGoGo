package tokenutil

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"os"
)

type TokenType int

const (
	ACCESS  TokenType = 0
	REFRESH TokenType = 1
	CSRF    TokenType = 2
)

func CreateToken(userID uuid.UUID, tokenType TokenType) (string, error) {
	switch tokenType {
	case ACCESS:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": userID,
		})
		return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	case REFRESH:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": userID,
		})
		return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	case CSRF:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": userID,
		})
		return token.SignedString([]byte("SECRET"))
	default:
		return "", fmt.Errorf("not valid token type")
	}
}

func ParseToken(tokenString string, tokenType TokenType) (*jwt.Token, error) {
	switch tokenType {
	case ACCESS:
		return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
	case REFRESH:
		return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("REFRESH_SECRET")), nil
		})
	case CSRF:
		return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("SECRET"), nil
		})
	default:
		return nil, fmt.Errorf("not valid token type")
	}
}

func CreateRedisKeyForToken(uuid string, token string) string {
	return fmt.Sprintf("%s:%s", uuid, token)
}
