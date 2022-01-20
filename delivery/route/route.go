package route

import (
	_authController "tixid/delivery/controller/authController"
	_userController "tixid/delivery/controller/userController"
	_middleware "tixid/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *_userController.UserController, ac *_authController.AuthController) {
	// route login
	e.POST("/login", ac.Login())

	// route user
	e.POST("/users", uc.Create())
	e.GET("/users", uc.GetAll())
	e.GET("/users/:id", uc.GetById())
	e.PUT("/users/:id", uc.Update(), _middleware.JWTMiddleware())
	e.DELETE("/users/:id", uc.Delete(), _middleware.JWTMiddleware())
}
