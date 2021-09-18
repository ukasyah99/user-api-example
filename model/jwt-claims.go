package model

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}
