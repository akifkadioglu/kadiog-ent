package authcontroller

import (
	"context"
	"net/http"
	"time"

	variables "setup/controllers/Variables"
	"setup/database"
	"setup/ent/user"
	"setup/env"
	"setup/helpers"
	"setup/localization"
	models "setup/models"
	language "setup/resources/Language"

	"entgo.io/ent/dialect/sql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {

	var input LoginInput
	db := database.DBManager()
	c.Bind(&input)

	if err := helpers.Validate(c, input); err != nil {
		return err
	}
	query := db.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.EQ(user.FieldEmail, input.Email))
	})
	user, err := query.First(context.Background())
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.PASSWORD_AND_EMAIL_DONT_MATCH, c),
		})
	}

	if err := helpers.CompareHash(user.Password, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.PASSWORD_AND_EMAIL_DONT_MATCH, c),
		})
	}

	var claims = &models.JwtCustomClaims{
		UserId: user.ID.String(),
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(), // For 1 year permission
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(env.Getenv(env.APP_KEY)))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		variables.TOKEN: t,
	})
}
