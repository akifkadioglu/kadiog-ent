package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	jwt.StandardClaims
}

func User(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
