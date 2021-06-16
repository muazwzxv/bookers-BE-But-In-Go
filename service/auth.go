package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/muazwzxv/bookers/m/config"
)

type JwtWrapper struct {
	SecretKey    string
	Issuer       string
	ExpiredHours int64
}

type JwtClaim struct {
	Email string
	Id    uint64
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(email string, id uint64) (signed string, err error) {
	claims := &JwtClaim{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpiredHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) VerifyToken(token string) error {
	if token != "" {
		return errors.New("token is null")
	}

	validate, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetJWTSecret()), nil
	})
	if err != nil {
		return err
	}

	if _, ok := validate.Claims.(jwt.MapClaims); !ok && !validate.Valid {
		return err
	}
	return nil
}
