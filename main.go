package main

import (
	"fmt"
	"log"
	"net/http"
	_config "tixid/config"
	_authController "tixid/delivery/controller/authController"
	_userController "tixid/delivery/controller/userController"
	_middleware "tixid/delivery/middleware"
	_route "tixid/delivery/route"
	_authRepository "tixid/repository/authRepository"
	_userRepository "tixid/repository/userRepository"
	_utility "tixid/utility"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := _config.GetConfig()
	db := _utility.InitDB(config)
	userRepo := _userRepository.New(db)
	userController := _userController.New(userRepo)

	authRepo := _authRepository.New(db)
	authController := _authController.New(*authRepo)

	e := echo.New()
	e.Pre(_middleware.CustomLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	_route.RegisterPath(e, userController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
