package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
)

func CreateToken(payload jwt.MapClaims, secret string) (string, error) {
	var err error
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetJwtToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateJwtToken(token, secret string) (jwt.MapClaims, error) {
	jwtToken, err := GetJwtToken(token, secret)
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error getting claims")
	}
	return claims, nil
}

func ValidateAuthHeader(header, secret string) (jwt.MapClaims, error) {
	splitted := strings.Split(header, " ")
	if len(splitted) != 2 {
		return nil, fmt.Errorf("wrong header")
	}
	return ValidateJwtToken(splitted[1], secret)
}

func HashPassword(plainText string) string {
	h := sha256.New()
	h.Write([]byte(plainText))
	return fmt.Sprintf("%x", h.Sum(nil))
}
