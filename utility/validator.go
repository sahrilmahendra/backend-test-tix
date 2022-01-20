package utility

import (
	_entity "tixid/entity"

	validator "github.com/go-playground/validator/v10"
)

type UserValidator struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Address  string `validate:"required"`
	Password string `validate:"required"`
}

func UserValidate(user _entity.User) error {
	v := validator.New()
	userValidate := UserValidator{
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password,
	}
	if err := v.Struct(userValidate); err != nil {
		return err
	}
	return nil
}

type LoginValidator struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func LoginValidate(login _entity.User) error {
	v := validator.New()
	loginValidate := LoginValidator{
		Username: login.Username,
		Password: login.Password,
	}
	if err := v.Struct(loginValidate); err != nil {
		return err
	}
	return nil
}
