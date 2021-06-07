package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (j *JwtWrapper) ValidateToken()
