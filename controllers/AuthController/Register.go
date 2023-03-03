package authcontroller

import (
	"context"
	"net/http"
	variables "setup/controllers/Variables"
	"setup/database"
	"setup/ent"
	"setup/ent/user"
	"setup/helpers"
	"setup/localization"
	language "setup/resources/Language"

	//emails "setup/resources/Views/Emails"

	"entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
)

type RegisterInput struct {
	Email                string `json:"email" validate:"required,email"`
	Name                 string `json:"name" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func Register(c echo.Context) error {
	db := database.DBManager()
	var input RegisterInput
	var err error
	var result *ent.User

	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := helpers.Validate(c, input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: err.Error(),
		})
	}
	query := db.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.EQ(user.FieldEmail, input.Email))
	})

	if users, err := query.All(context.Background()); len(users) > 0 || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.EMAIL_ALREADY_EXIST, c),
		})
	}

	if input.Password != input.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, map[string]string{
			variables.MESSAGE: localization.TR(language.PASSWORDS_DONT_MATCH, c),
		})
	}

	if result, err = db.User.Create().
		SetPassword(helpers.Hash(input.Password)).
		SetEmail(input.Email).
		SetName(input.Name).
		Save(context.Background()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			variables.MESSAGE: localization.TR(language.SOMETHING_WENT_WRONG, c),
			variables.TOKEN:   err.Error(),
		})
	}

	/* if err := helpers.SendEmail(user.Email, localization.TR(language.NEW_ACCOUNT, c), emails.Register(user.Name, c)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			variables.MESSAGE: localization.TR(language.SOMETHING_WENT_WRONG, c),
		})
	} */

	return c.JSON(http.StatusOK, result)
}
