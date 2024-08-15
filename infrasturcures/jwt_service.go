package infrastructure

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	JwtSecret []byte
}

func (j *JwtService) GenerateToken(email string, isSeller bool, expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  isSeller,
		"exp":   expirationTime,
	})
	jwtToken, e := token.SignedString(j.JwtSecret)

	if e != nil {
		return "", errors.New("can't sign token")
	}

	return jwtToken, nil
}

// validate token
func (j *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.JwtSecret, nil
	})

	if err != nil || !jwtoken.Valid {
		return nil, errors.New("invalid JWT")
	}

	return jwtoken, err
}

func (j *JwtService) ValidateAdmin(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	return ok && claims["is_admin"].(bool)
}