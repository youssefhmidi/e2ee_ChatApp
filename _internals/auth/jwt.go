package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

func CreateAccessToken(user models.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Duration(expiry) * time.Hour)
	claim := models.JwtCustomClaims{
		Id:   int(user.ID),
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user models.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Duration(expiry) * time.Hour)
	claim := models.JwtCustomClaims{
		Id:   int(user.ID),
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))
}

// returns a boolean and an error
func ValidateToken(token string, secret string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot indentify the signing method used")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// returns an id with the uint type and an error
func GetIdFromToken(token string, secret string) (uint, error) {
	Token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot indentify the signing method used")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	Raw, ok := Token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}
	return uint(Raw["id"].(float64)), nil
}
