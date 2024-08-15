package infrastructures

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	JwtSecret []byte
}

func (j *JwtService) GenerateToken(email string, isSeller bool) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": 	email,
		"seller":  	isSeller,
		"refresh": 	false,
		"exp":   	time.Now().Add(5 * time.Minute).Unix(),
	})
	accessToken, e := token.SignedString(j.JwtSecret)

	if e != nil {
		return "", "", errors.New("can't sign token")
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": 	email,
		"seller":  	isSeller,
		"refresh": 	true,
		"exp":   	time.Now().Add(2160 * time.Hour).Unix(),
	})
	refreshToken, e := token.SignedString(j.JwtSecret)

	if e != nil {
		return "", "", errors.New("can't sign token")
	}

	return accessToken, refreshToken, nil
}

// validate token
func (j *JwtService) ValidateToken(token string, isRefresh bool) (string, bool, error) {
	jwtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.JwtSecret, nil
	})

	if err != nil || !jwtoken.Valid {
		return "", false, errors.New("invalid JWT")
	}

	claims, ok := jwtoken.Claims.(jwt.MapClaims)
	
	if !ok || (!claims["refresh"].(bool) && isRefresh) || (claims["refresh"].(bool) && !isRefresh) {
		return "", false, errors.New("invalid JWT")
	}

	return claims["email"].(string) ,claims["seller"].(bool), err
}


func (j *JwtService) GenerateActivationToken(email string, isSeller bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": 	email,
		"seller":  	isSeller,
		"refresh": 	false,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})
	activationToken, e := token.SignedString(j.JwtSecret)

	if e != nil {
		return "", errors.New("can't sign token")
	}

	return activationToken, nil
}