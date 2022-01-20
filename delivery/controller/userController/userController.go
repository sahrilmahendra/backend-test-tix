package userController

import (
	"net/http"
	"strconv"
	_controller "tixid/delivery/controller"
	_middleware "tixid/delivery/middleware"
	_entity "tixid/entity"
	_userRepository "tixid/repository/userRepository"
	_utility "tixid/utility"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userRepo _userRepository.User
}

func New(user _userRepository.User) *UserController {
	return &UserController{userRepo: user}
}

func (uc UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := _entity.User{}

		if err := c.Bind(&newUser); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}

		if err := _utility.UserValidate(newUser); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}

		newUser.Password, _ = _utility.HashPassword(newUser.Password)
		if _, err := uc.userRepo.Create(newUser); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "username or email already exist")
		}

		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.userRepo.GetAll()
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get user")
		} else if len(users) == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "user not found")
		}
		var userResponse []UserResponseFormat
		for _, value := range users {
			userResponse = append(userResponse, FormattingUserResponse(value))
		}
		return _controller.SuccessWithDataResponse(c, "success operation", userResponse)
	}
}

func (uc UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid user id")
		}

		user, lenUser, err := uc.userRepo.GetById(id)
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get user")
		} else if lenUser == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "user not found")
		}

		userResponse := FormattingUserResponse(user)

		return _controller.SuccessWithDataResponse(c, "success operation", userResponse)
	}
}

func (uc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		if valid := _middleware.ValidateToken(c); !valid {
			return _controller.ErrorResponse(c, http.StatusUnauthorized, "access forbidden")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}
		_, lenUser, _ := uc.userRepo.GetById(id)
		if lenUser == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "user not found")
		}
		updateUser := _entity.User{}
		if err := c.Bind(&updateUser); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}
		if err := _utility.UserValidate(updateUser); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}
		updateUser.Password, _ = _utility.HashPassword(updateUser.Password)
		if _, err = uc.userRepo.Update(id, updateUser); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "username or email already exist")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

func (uc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		if valid := _middleware.ValidateToken(c); !valid {
			return _controller.ErrorResponse(c, http.StatusUnauthorized, "access forbidden")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}
		_, lenUser, _ := uc.userRepo.GetById(id)
		if lenUser == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "user not found")
		}
		if err = uc.userRepo.Delete(id); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to delete user")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}
