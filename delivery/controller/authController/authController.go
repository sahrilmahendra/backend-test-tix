package authcontroller

import (
	"net/http"
	_controller "tixid/delivery/controller"
	_entity "tixid/entity"
	_authRepository "tixid/repository/authRepository"
	_utility "tixid/utility"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authRepo _authRepository.AuthRepository
}

func New(auth _authRepository.AuthRepository) *AuthController {
	return &AuthController{authRepo: auth}
}

func (ac AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := _entity.User{}

		if err := c.Bind(&user); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}

		if err := _utility.LoginValidate(user); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}

		token, err := ac.authRepo.Login(user)
		if err != nil || token == "" {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "login failed")
		}

		return _controller.SuccessWithDataResponse(c, "login success", token)
	}
}
