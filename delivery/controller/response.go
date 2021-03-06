package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponseNonData struct {
	Meta struct {
		Code    int    `json:"Code"`
		Message string `json:"Message"`
	} `json:"Meta"`
}

type BaseResponseWithData struct {
	Meta struct {
		Code    int    `json:"Code"`
		Message string `json:"Message"`
	} `json:"Meta"`
	Data interface{} `json:"Data"`
}

func SuccessNonDataResponse(c echo.Context, message string) error {
	response := BaseResponseNonData{}
	response.Meta.Code = http.StatusOK
	response.Meta.Message = message

	return c.JSON(http.StatusOK, response)
}

func SuccessWithDataResponse(c echo.Context, message string, data interface{}) error {
	response := BaseResponseWithData{}
	response.Meta.Code = http.StatusOK
	response.Meta.Message = message
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, message string) error {
	response := BaseResponseNonData{}
	response.Meta.Code = status
	response.Meta.Message = message

	return c.JSON(status, response)
}
