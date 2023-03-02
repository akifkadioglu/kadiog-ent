package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	jwt.StandardClaims
}
