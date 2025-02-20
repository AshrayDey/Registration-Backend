package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GenerateJWT(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected Signing Method")
		}

		return secretKey, nil
	})
	if err != nil {
		return 0, errors.New("Could not parse the token")
	}

	tokenIsValied := parsedToken.Valid
	if !tokenIsValied {
		return 0, errors.New("Token not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid claims")
	}
	// email :=claims["email"].(string)
	// userId := claims["userId"].(int64)
	return claims["userId"].(int64), nil
}
