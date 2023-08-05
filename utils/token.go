package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {

	token := jwt.New(jwt.SigningMethodES256)

	timenow := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = timenow.Add(ttl).Unix()
	claim["iat"] = timenow.Unix()
	claim["nbf"] = timenow.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("failed generating JWT token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(token, signedJWTKey string) (interface{}, error) {

	tokenParse, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token with %v", err)
	}
	claims, ok := tokenParse.Claims.(jwt.MapClaims)
	if !ok || !tokenParse.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}
