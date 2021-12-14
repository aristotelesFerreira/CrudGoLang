package validators

import (
	"errors"
	"strings"

	"github.com/aristotelesFerreira/crud_go_lang/api/models"
	"github.com/badoux/checkmail"
)

func UserValidate(u models.User, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required nickname")
		}
		if u.Password == "" {
			return errors.New("Required password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required nickname")
		}
		if u.Password == "" {
			return errors.New("Required password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	}
}
