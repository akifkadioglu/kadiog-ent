package authcontroller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterInput struct {
	Email                string `json:"email" validate:"required,email"`
	Name                 string `json:"name" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	Username             string `json:"username" validate:"required"`
}

func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register")
}
