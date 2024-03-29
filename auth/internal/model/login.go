package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginIn struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,password"`
}

func (in LoginIn) Validate() error {
	err := validation.ValidateStruct(&in,
		validation.Field(&in.Email, validation.Required, validation.Length(EmailMinLength, EmailMaxLength), is.Email),
		validation.Field(&in.Password, validation.Required, validation.Length(PasswordMinLength, PasswordMaxLength)),
	)
	if err != nil {
		return err // fmt.Errorf("validation error: %w", err)
	}

	return nil
}

type LoginOut struct {
	Token string
}
