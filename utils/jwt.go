package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const key = "24d75f9989d1426ad0f659cae445ce6ca36d950866e5a9440786079433745132"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(key))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}
		return []byte(key), nil
	})

	if err != nil {
		return errors.New("Could not parse token: " + err.Error() + ".")
	}

	if !parsedToken.Valid {
		return errors.New("Invalid token.")
	}

	/*
		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		if !ok {
			return errors.New("Invalid token claims.")
		}
		email := claims["email"].(string)
		userId := claims["userId"].(int64)
	*/
	return nil
}
